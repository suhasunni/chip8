package cpu

import (
	"math/rand"
)

// to do: reorder the functions based on how they are decoded

// Store 2 byte word into byte indexed memory (Big Endian)
func (c *CPU) writeWordToMemory(imm uint16, addr uint8) {
	c.memory[addr] = uint8(imm >> 8)
	c.memory[addr+1] = uint8(imm)
}

// 00E0 - Clear display
func (c *CPU) clear() {

}

// 00EE - Return from a subroutine
func (c *CPU) ret() {
	c.pc = uint16(c.memory[c.sp])
	c.sp--
}

// 1nnn - Jump to nnn
func (c *CPU) jump(imm uint16) {
	c.pc = uint16(imm)
}

// 2nnn - Call subroutine at nnn
func (c *CPU) call(imm uint16) {
	// TO DO: add a check for stack overflow here
	c.sp++
	c.writeWordToMemory(imm, c.sp)
	c.pc = imm
}

// 3xkk - Skip next instruction if Vx == kk
func (c *CPU) skipIfEqualImm(reg uint8, imm uint8) {
	if c.registers[reg] == imm {
		c.pc += 2
	}
}

// 4xkk - Skip next instruction if Vx != kk
func (c *CPU) skipIfNotEqualImm(reg uint8, imm uint8) {
	if c.registers[reg] != imm {
		c.pc += 2
	}
}

// 5xy0 - Skip next instruction if Vx == Vy
func (c *CPU) skipIfEqual(reg1 uint8, reg2 uint8) {
	if c.registers[reg1] == c.registers[reg2] {
		c.pc += 2
	}
}

// 9xy0 - Skip instruction if Vx == Vy
func (c *CPU) skipIfNotEqual(reg1 uint8, reg2 uint8) {
	if c.registers[reg1] != c.registers[reg2] {
		c.pc += 2
	}
}

// 6xkk - Set Vx = kk
func (c *CPU) setReg(reg uint8, imm uint8) {
	c.registers[reg] = imm
}

// 7xkk - Set Vx == Vx + imm
func (c *CPU) addImm(reg uint8, imm uint8) {
	c.registers[reg] += imm
}

// 8xy0 - Set Vx = Vy
func (c *CPU) assign(reg1 uint8, reg2 uint8) {
	c.registers[reg1] = c.registers[reg2]
}

// 8xy1 - Set Vx = Vx OR Vy
func (c *CPU) or(reg1 uint8, reg2 uint8) {
	c.registers[reg1] |= c.registers[reg2]
}

// 8xy2 - Set Vx = Vx AND Vy
func (c *CPU) and(reg1 uint8, reg2 uint8) {
	c.registers[reg1] &= c.registers[reg2]
}

// 8xy3 - Set Vx = Vx XOR Vy
func (c *CPU) xor(reg1 uint8, reg2 uint8) {
	c.registers[reg1] ^= c.registers[reg2]
}

// 8xy4 - Set Vx = Vx + Vy, set Vf = carry
func (c *CPU) add(reg1 uint8, reg2 uint8) {
	val1, val2 := uint16(c.registers[reg1]), uint16(c.registers[reg2])
	result := val1 + val2
	if result > 255 {
		c.registers[0xf] = 1
	} else {
		c.registers[0xf] = 0
	}
	c.registers[reg1] = uint8(result)
}

// 8xy5 - Set Vx = Vx - Vy, set Vf = no borrow
func (c *CPU) sub(reg1 uint8, reg2 uint8) {
	if c.registers[reg1] < c.registers[reg2] {
		c.registers[0xf] = 0
	} else {
		c.registers[0xf] = 1
	}
	c.registers[reg1] -= c.registers[reg2]
}

// 8xy6 - Shift Vx right by 1 bit, Set Vf = remainder
func (c *CPU) shiftRight(reg uint8) {
	if reg%2 == 0 {
		c.registers[0xf] = 0
	} else {
		c.registers[0xf] = 1
	}
	c.registers[reg] >>= 1
}

// 8xy7 - Set Vx = Vy - Vx, set Vx = no borrow
func (c *CPU) subn(reg1 uint8, reg2 uint8) {
	if c.registers[reg2] < c.registers[reg1] {
		c.registers[0xf] = 0
	} else {
		c.registers[0xf] = 1
	}
	c.registers[reg1] = c.registers[reg2] - c.registers[reg1]
}

// 8xyE - Shift Vx left by 1 bit, set Vf = (MSB[reg1] == 1)
func (c *CPU) shiftLeft(reg uint8) {
	if c.registers[reg]>>7 == 1 {
		c.registers[0xf] = 1
	} else {
		c.registers[0xf] = 0
	}
	c.registers[reg] <<= 1
}

// Annn - Set ir = nnn
func (c *CPU) setIR(imm uint16) {
	c.ir = imm
}

// Bnnn - Set pc = nnn + V0
func (c *CPU) setPC(imm uint16) {
	c.pc = uint16(c.registers[0]) + imm
}

// Cxkk - Set Vx = random byte AND kk
func (c *CPU) setRand(reg uint8, imm uint8) {
	c.registers[reg] = uint8(rand.Intn(256)) & imm
}
