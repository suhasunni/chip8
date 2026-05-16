package cpu

import (
	"math/rand"
)

// to do: reorder the functions based on how they are decoded

// Store 2 byte word into byte-indexed memory (Big Endian)
func (c *CPU) writeWordToMemory(imm uint16, addr uint8) {
	c.memory[addr] = uint8(imm >> 8)
	c.memory[addr+1] = uint8(imm)
}

// 00E0 - Clear display
func (c *CPU) clear() {
	for i := range 64 {
		for j := range 32 {
			c.Buffer[i][j] = false
		}
	}
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

// 9xy0 - Skip instruction if Vx != Vy
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
	if c.registers[reg]%2 == 0 {
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

// Dxyn - Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision.
func (c *CPU) display(reg1 uint8, reg2 uint8, n uint8) {
	// to do: implement
	return
}

// Ex9E - Skip instruction if value at Vx is pressed
func (c *CPU) skipIfPressed(reg uint8) {
	// to do: implement
	return
}

// ExA1- Skip instruction if value at Vx is not pressed
func (c *CPU) skipIfNotPressed(reg uint8) {
	// to do: implement
	return
}

// Fx07 - Set Vx = delayTimer value
func (c *CPU) delayValue(reg uint8) {
	c.registers[reg] = c.delayTimer
}

// Fx0A - Store value of next key press in Vx
func (c *CPU) nextKeyPress(reg uint8) {
	// to do: implement
	return
}

// Fx15 - Set delay timer to Vx
func (c *CPU) setDelay(reg uint8) {
	c.delayTimer = c.registers[reg]
}

// Fx18 - Set sound timer to Vx
func (c *CPU) setSound(reg uint8) {
	c.soundTimer = c.registers[reg]
}

// Fx1E - Set IR = IR + Vx
func (c *CPU) addIR(reg uint8) {
	c.ir += uint16(c.registers[reg])
}

// Fx29 - Set IR to sprite location of Vx
func (c *CPU) getSprite(reg uint8) {
	c.ir = fontOffset + (uint16(c.registers[reg]))*5
}

// Fx33 - Store decimal digits of Vx in I, I+1, and I+2
func (c *CPU) storeIR(reg uint8) {
	num := c.registers[reg]
	c.memory[c.ir] = num / 100
	num %= 100
	c.memory[c.ir+1] = num / 10
	c.memory[c.ir+2] = num % 10
}

// Fx55 - Store registers V0 to Vx in memory, starting at IR
func (c *CPU) storeRegisters(reg uint8) {
	for i := range reg + 1 {
		c.memory[c.ir+uint16(i)] = c.registers[i]
	}
}

// Fx65 - Load registers V0 to Vx from memory
func (c *CPU) loadRegisters(reg uint8) {
	for i := range reg + 1 {
		c.registers[i] = c.memory[c.ir+uint16(i)]
	}
}
