package main

import (
	"fmt"
	"tetris"
)

func main() {
	s := tetris.ReadFile()
	a := [][]string{}
	for i := 0; i < len(s); i++ {
		if len(s[i]) > 0 {
			a = append(a, []string{s[i]})
		}
	}

	for j := 0; j < len(a); j++ {
		for k := 0; k < len(a[j]); k++ {
			for l := 0; l < len(a[j][k]); l++ {
				fmt.Println(a[j][k][l])
			}
		}
	}
}
