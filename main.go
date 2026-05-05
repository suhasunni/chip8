package main

import (
	"fmt"
	"github.com/suhasunni/chip8/cpu"
)

func main() {
	core := cpu.NewCPU()
	fmt.Println(core.String())
}

/*
to do:
	start stack
*/
