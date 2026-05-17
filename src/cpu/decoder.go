package cpu

func (c *CPU) decodeAndExecute(instruction uint16) {
	reg1, reg2 := (instruction&0x0F00)>>8, (instruction&0x00F0)>>4
	switch (instruction & 0xF000) >> 12 {
	case 0x0:
		switch instruction & 0x00FF {
		case 0xE0:
			c.clear()
		case 0xEE:
			c.ret()
		}
	case 0x1:
		c.jump(instruction & 0x0FFF)
	case 0x2:
		c.call(instruction & 0x0FFF)
	case 0x3:
		c.skipIfEqualImm(reg1, byte(instruction&0x00FF))
	case 0x4:
		c.skipIfNotEqualImm(reg1, byte(instruction&0x00FF))
	case 0x5:
		c.skipIfEqual(reg1, reg2)
	case 0x6:
		c.setRegister(reg1, byte(instruction&0x00FF))
	case 0x7:
		c.addImm(reg1, byte(instruction&0x00FF))
	case 0x8:
		switch instruction & 0x00F {
		case 0x0:
			c.assign(reg1, reg2)
		case 0x1:
			c.or(reg1, reg2)
		case 0x2:
			c.and(reg1, reg2)
		case 0x3:
			c.xor(reg1, reg2)
		case 0x4:
			c.add(reg1, reg2)
		case 0x5:
			c.sub(reg1, reg2)
		case 0x6:
			c.shiftRight(reg1)
		case 0x7:
			c.subn(reg1, reg2)
		case 0xE:
			c.shiftLeft(reg1)
		}
	case 0x9:
		c.skipIfNotEqual(reg1, reg2)
	case 0xA:
		c.setIR(instruction & 0x0FFF)
	case 0xB:
		c.setPC(instruction & 0x0FFF)
	case 0xC:
		c.setRand(reg1, byte(instruction&0x00FF))
	case 0xD:
		c.display(reg1, reg2, byte(instruction&0x000F))
	case 0xE:
		switch instruction & 0x00FF {
		case 0x9E:
			c.skipIfPressed(reg1)
		case 0xA1:
			c.skipIfNotPressed(reg1)
		}
	case 0xF:
		switch instruction & 0x00FF {
		case 0x07:
			c.delayValue(reg1)
		case 0x0A:
			c.nextKeyPress(reg1)
		case 0x15:
			c.setDelay(reg1)
		case 0x18:
			c.setSound(reg1)
		case 0x1E:
			c.addIR(reg1)
		case 0x29:
			c.getSprite(reg1)
		case 0x33:
			c.storeIR(reg1)
		case 0x55:
			c.storeRegisters(reg1)
		case 0x65:
			c.loadRegisters(reg1)
		}
	}
}
