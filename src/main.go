package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
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
