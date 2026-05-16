package cpu

import (
	"fmt"
	"log"
	"os"
)

type CPU struct {
	memory     [4096]byte // main memory
	registers  [16]byte   // registers
	pc         uint16     // program counter
	sp         byte       // stack pointer
	ir         uint16     // index register
	delayTimer byte
	soundTimer byte
	Buffer     [64][32]bool // display
	stack      [16]uint16   // call stack
	keypad     [16]bool
}

const fontOffset uint16 = 0x50

func NewCPU() *CPU {
	// to do: fix this text
	log.Println("Creating CPU...")
	c := CPU{}

	log.Println("Loading Fonts...")
	c.loadFontsIntoMemory()
	log.Println("Fonts loaded into memory successfully.")

	log.Println("Loading ROM file...")
	c.loadROM()
	log.Println("ROM loaded in memory successfully.")

	return &c
}

func (c *CPU) DecrementTimers() {
	if c.delayTimer > 0 {
		c.delayTimer--
	}
	if c.soundTimer > 0 {
		c.soundTimer--
	}
}

// stack points to the next available spot on the stack, sp = [0, 15]
func (c *CPU) pushStack(imm uint16) error {
	if c.sp == 16 {
		return fmt.Errorf("Stack Overflow")
	}
	c.stack[c.sp] = imm
	c.sp++
	return nil
}

func (c *CPU) popStack() (uint16, error) {
	if c.sp == 0 {
		return 0, fmt.Errorf("Stack Underflow")
	}
	c.sp--
	return c.stack[c.sp], nil
}

func (c *CPU) Tick() {
	// fetch instruction
	instruction := uint16(c.memory[c.pc])<<8 | uint16(c.memory[c.pc+1])
	c.pc += 2

	// decode/excecute
	var mask byte = 0xF
	var nibble1 byte = byte((instruction & (uint16(mask) << 12)) >> 12)
	var nibble2 byte = byte((instruction & (uint16(mask) << 8)) >> 8)
	var nibble3 byte = byte((instruction & (uint16(mask) << 4)) >> 4)
	var nibble4 byte = byte(instruction & uint16(mask))

	switch nibble1 {
	case 0:
		c.clear()
	case 1:
		c.jump(combineBytes(nibble2, nibble3, nibble4))
	case 6:
		c.setRegister(nibble2, (nibble3<<4)|nibble4)
	case 0xA:
		c.setIR(combineBytes(nibble2, nibble3, nibble4))
	case 7:
		c.addImm(nibble2, (nibble3<<4)|nibble4)
	case 0xD:
		c.display(nibble2, nibble3, nibble4)
	}
}

func (c *CPU) loadROM() {
	data, err := os.ReadFile("test.ch8")
	if err != nil {
		log.Fatal(err)
	}
	// start loading data at 0x200
	for i := range len(data) {
		c.memory[0x200+uint16(i)] = data[i]
	}
	c.pc = 0x200
}

// helper to make three bytes a uint16
func combineBytes(b1 byte, b2 byte, b3 byte) uint16 {
	return (uint16(b1) << 8) | (uint16(b2) << 4) | uint16(b3)
}
