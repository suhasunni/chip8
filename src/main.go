package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/suhasunni/chip8/src/cpu"
	"log"
)

func main() {
	core := cpu.NewCPU()
	fmt.Println(core.String())

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Chip-8 Emulator")
	if err := ebiten.RunGame(NewDisplay()); err != nil {
		log.Fatal(err)
	}
}

/*
to do:
 - add logging stages of loading
*/
