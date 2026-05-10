package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/suhasunni/chip8/src/cpu"
)

type Display struct {
	cpu		cpu.CPU
	buffer [64][32]bool
}

func (d *Display) Update() error {
	// check which keys are pressed use ebiten.IsKeyPressed, then pass that input to the cpu
	return nil
}

func (d *Display) Draw(screen *ebiten.Image) {
	// create an image by iterating over 64x32 buffer, then draw that image to the screen
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (d *Display) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 64, 32
}
