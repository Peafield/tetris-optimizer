package tetris

import "fmt"

func PrintBoard(s string) {
	r := []rune(s)
	row := []string{}
	h := 0
	a := 30
	for i := 0; i < len(r); i++ {
		if r[i] == 10 {
			PrintRow(row)
			row = []string{}
		} else if r[i] == '#' {
			h += 1
			if h == 5 {
				a += 1
				h = 1
			}
			row = append(row, string(r[i]+rune(a)))
		}
	}
}

func PrintRow(r []string) {
	for i, v := range r {
		fmt.Print(v)
		if i == len(r)-1 {
			fmt.Println()
		}
	}
}
