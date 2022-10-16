package main

import (
	"fmt"
	"tetris"
)

func main() {
	s := tetris.ReadFile()
	a := []string{}
	b := []string{}
	str := ""
	r := 30
	for i := 0; i < len(s); i++ {
		if len(s[i]) > 0 {
			a = append(a, string(s[i]))
		}
	}
	fmt.Println(a)
	for row := 0; row < len(a); row++ {
		for col := 0; col < len(a[row]); col++ {
			if col == len(a[row])-1 || row == len(a)-1 {
				fmt.Println(str)
				b = append(b, str)
				str = ""
			}
			if a[row][col] == '#' {
				str += string(int(a[row][col]) + r)
			}
		}
	}
	for n := 0; n < len(b); n++ {
		fmt.Println([]byte(b[n]))
	}
}
