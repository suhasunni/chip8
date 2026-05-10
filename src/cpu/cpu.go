package cpu

type CPU struct {
	memory     [4096]uint8 // main memory
	registers  [16]uint8   // registers
	pc         uint16      // program counter
	sp         uint8       // stack pointer
	ir         uint16      // index register
	delayTimer uint8
	soundTimer uint8
	font       [16][5]uint16
	Buffer     [64][32]bool
}

const fontOffset uint16 = 0x50

func NewCPU() *CPU {
	c := CPU{}
	return &c
}

func (c CPU) String() string {
	return "cpu made"
}

func (c *CPU) DecrementTimers() {
	if c.delayTimer > 0 {
		c.delayTimer--
	}
	if c.soundTimer > 0 {
		c.soundTimer--
	}
}
