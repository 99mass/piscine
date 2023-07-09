package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n > 0 {
		var c []int
		a := 0

		b := 0
		var d int
		for n != 0 {
			a = n % 10
			n /= 10
			c = append(c, a)
		}

		for count := range c {
			b = count + 1
		}
		for i := 0; i < b-1; i++ {
			for j := 0; j < b-i-1; j++ {
				if c[j] > c[j+1] {
					d = c[j]
					c[j] = c[j+1]
					c[j+1] = d
				}
			}
		}
		for i := 0; i < b; i++ {
			z01.PrintRune(rune(c[i] + 48))
		}
	}
}
