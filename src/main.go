package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"os"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Chip-8 Emulator")
	if err := ebiten.RunGame(NewDisplay(os.Args[1])); err != nil {
		log.Fatal(err)
	}
}
