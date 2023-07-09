package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	arguments2 := os.Args

	for i := 1; i < len(arguments); i++ {
		for _, r := range arguments2[i] {
			z01.PrintRune(r)
		}
		z01.PrintRune(rune('\n'))
	}
}
