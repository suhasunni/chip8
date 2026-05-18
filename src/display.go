package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/suhasunni/chip8/src/cpu"
)

type Display struct {
	cpu    *cpu.CPU
	keyMap map[ebiten.Key]int
}

func NewDisplay() *Display {
	d := &Display{cpu: cpu.NewCPU()}
	d.keyMap = map[ebiten.Key]int{
		ebiten.Key1: 1,
		ebiten.Key2: 2,
		ebiten.Key3: 3,
		ebiten.Key4: 12,
		ebiten.KeyQ: 4,
		ebiten.KeyW: 5,
		ebiten.KeyE: 6,
		ebiten.KeyR: 13,
		ebiten.KeyA: 8,
		ebiten.KeyS: 9,
		ebiten.KeyD: 10,
		ebiten.KeyF: 14,
		ebiten.KeyZ: 10,
		ebiten.KeyX: 0,
		ebiten.KeyC: 11,
		ebiten.KeyV: 15,
	}
	return d
}

func (d *Display) Update() error {
	// check for keyboard input
	for key, val := range d.keyMap {
		if ebiten.IsKeyPressed(key) {
			d.cpu.Keypad[val] = true
		} else {
			d.cpu.Keypad[val] = false
		}
	}

	d.cpu.DecrementTimers()
	d.cpu.Tick()
	d.cpu.Tick()
	d.cpu.Tick()
	d.cpu.Tick()
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
