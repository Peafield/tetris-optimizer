package main

import (
	"tetris"
)

func main() {
	s := tetris.ReadFile()
	tetris.PrintBoard(s)
}
