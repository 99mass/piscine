// package main

// 00 01, 00 02, 00 03, ..., 00 98, 00 99, 01 02, 01 03, ..., 97 98, 97 99, 98 99
package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i = i + 1 {
		for j := '0'; j <= '9'; j = j + 1 {
			m := j + 1
			for k := i; k <= '9'; k = k + 1 {
				for ; m <= '9'; m = m + 1 {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(m)
					if i < '9' || j < '8' || k < '9' || m < '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				m = '0'
			}
		}
	}
	z01.PrintRune('\n')
}

/*func main() {
	PrintComb2()
}*/
