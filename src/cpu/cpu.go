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
}

func NewCPU() *CPU {
	c := CPU{}
	return &c
}

func (c CPU) String() string {
	return "cpu made"
}

func (c *CPU) decrementDelayTimer() {
	if c.delayTimer > 0 {
		c.delayTimer--
	}
}

func (c *CPU) decrementSoundTimer() {
	if c.soundTimer > 0 {
		c.soundTimer--
	}
}
