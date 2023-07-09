package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := []rune(os.Args[0])

	for i := 2; i < len(arguments); i++ {
		z01.PrintRune(arguments[i])
	}
	z01.PrintRune('\n')
}
