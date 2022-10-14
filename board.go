package tetris

import "fmt"

func PrintBoard(s string) {
	r := []rune(s)
	row := []string{}
	for i := 0; i < len(r); i++ {
		if r[i] == 10 {
			PrintRow(row)
			row = []string{}
		} else if r[i] == '#' {
			row = append(row, string(r[i]))
		}
	}
	// PrintRow(row)
}

func PrintRow(r []string) {
	for i := 0; i < len(r); i++ {
		fmt.Println(r[i])
	}
}
