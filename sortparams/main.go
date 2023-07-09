package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for i := 1; i < len(arg); i++ {
		for k := 1; k < len(arg); k++ {
			if arg[k] > arg[i] {
				arg[i], arg[k] = arg[k], arg[i]
			}
		}
	}
	for k := 1; k < len(arg); k++ {
		for _, l := range arg[k] {
			z01.PrintRune(l)
		}
		z01.PrintRune('\n')
	}
}
