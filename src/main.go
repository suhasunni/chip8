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
 - make the terminal run command take the file name as an argument
*/
