package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	arg := os.Args
	for i := (len(argument)) - 1; i >= 1; i-- {
		for _, v := range arg[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
