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
 - make helpers for decoding (e.g 3 bytes to a uint16)
 - in alu make comments align with func signature or instructions (preferable)
	- maybe add a comment on the top clarifying Vx = reg[x]
	- also make func names consistient (e.g reg vs register)
 - add logging stages of loading
 - make the terminal run command take the file name as an argument
*/
