package piscine

func RecursiveFactorial(nb int) int {
	a := 1
	if nb < 0 || nb > 20 {
		return 0
	} else if nb == 0 {
		return 1
	}
	a = nb * RecursiveFactorial(nb-1)

	return a
}
