package cpu

import (
	"fmt"
	"log"
	"os"
)

type CPU struct {
	memory           [4096]byte // main memory
	registers        [16]byte   // registers
	pc               uint16     // program counter
	sp               byte       // stack pointer
	ir               uint16     // index register
	delayTimer       byte
	soundTimer       byte
	Buffer           [64][32]bool // display
	stack            [16]uint16   // call stack
	keypad           [16]bool
	waitingForKey    bool
	waitingForKeyReg uint16
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
	if c.waitingForKey {
		c.processInput()
		return
	}
	// fetch instruction
	instruction := uint16(c.memory[c.pc])<<8 | uint16(c.memory[c.pc+1])
	c.pc += 2

	c.decodeAndExecute(instruction)
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

func (c *CPU) processInput() {
	for i := range 16 {
		if c.keypad[i] {
			c.registers[c.waitingForKeyReg] = byte(i)
			c.waitingForKey = false
			break
		}
	}
}
