package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/suhasunni/chip8/src/cpu"
	"image/color"
)

type Display struct {
	cpu *cpu.CPU
}

func NewDisplay() *Display {
	d := &Display{cpu: cpu.NewCPU()}
	return d
}

func (d *Display) Update() error {
	// check which keys are pressed use ebiten.IsKeyPressed, then pass that input to the cpu, if ANY key is pressed, update c.waitingForKey
	d.cpu.Tick()
	return nil
}

// create an image by iterating over 64x32 buffer, then draw that image to the screen
func (d *Display) Draw(screen *ebiten.Image) {

	for i := range 64 {
		for j := range 32 {
			if d.cpu.Buffer[i][j] {
				screen.Set(i, j, color.White)
			} else {
				screen.Set(i, j, color.Black)
			}
		}
	}

}

func (d *Display) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 64, 32
}

func (d *Display) clearBuffer(c *cpu.CPU) {
	for i := range 64 {
		for j := range 32 {
			d.cpu.Buffer[i][j] = false
		}
	}
}
