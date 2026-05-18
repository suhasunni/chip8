# CHIP-8 Interpreter

A CHIP-8 interpreter written in Go.

https://github.com/user-attachments/assets/a5d1042e-93c5-4dd6-b40e-9ef52b5cf5bb

  <em>A rock-paper-scissors game by [SystemLogoff](https://systemlogoff.com/) running on the Chip-8 Interpreter</em>

## Project Description

CHIP-8 is an interpreted programming language developed in the 1970s for microcomputers. It was designed to make game development easier on limited hardware and quickly became a popular platform for simple arcade-style games.

This project implements the core CHIP-8 system, including:

* Memory and register management
* Instruction decoding and execution
* Graphics rendering
* Timers
* Keypad input handling
* ROM loading and execution
* Call stack

The interpreter loads CHIP-8 ROM files and executes them through a fetch-decode-execute cycle, updating graphics and input state in real time. Classic games such as Pong, Space Invaders, Tetris, and Breakout can be played through compatible ROMs.

## Requirements

Before building the project, ensure the following dependencies are installed:

* Go 1.25.4 or newer
* Ebiten v2.9.9

This project uses the following primary library:

* `github.com/hajimehoshi/ebiten/v2`

Additional indirect dependencies are managed automatically through Go modules.

## Building

Clone the repository:

```bash
git clone https://github.com/suhasunni/chip8.git
cd chip8
```

Build the project:

```bash
go build -o chip8 .
```

## Running

The interpreter expects a CHIP-8 ROM file as a command-line argument.

Example:

```bash
./chip8 path/to/rom.ch8
```

Once launched, the interpreter will load the ROM into memory and begin execution. Games such as Tetris, Pong, and Space Invaders are included in `testgames/`.

## Controls

CHIP-8 programs use a 16-key hexadecimal keypad.

Original COSMAC VIP keypad:
```text
1 2 3 C
4 5 6 D
7 8 9 E
A 0 B F
```

Keyboard mapping:
```text
1 2 3 4
Q W E R
A S D F
Z X C V
```


## Resources

This project was inspired by Tobias Langhoff’s excellent high-level, code-free blog post on writing a CHIP-8 emulator: https://tobiasvl.github.io/blog/write-a-chip-8-emulator/

The CHIP-8 technical reference used for implementation details: http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#keyboard
