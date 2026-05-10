package main

import (
	"fmt"
	"log"
	"github.com/suhasunni/chip8/src/cpu"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	core := cpu.NewCPU()
	fmt.Println(core.String())
	
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Chip-8 Emulator")
	if err := ebiten.RunGame(&Display{}); err != nil {
		log.Fatal(err)
	}
}

/*
to do:
	start stack
*/
