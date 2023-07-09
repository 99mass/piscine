package piscine

func IterativeFactorial(nb int) int {
	a := 1
	if nb < 0 || nb > 20 {
		return 0
	} else if nb == 0 {
		return 1
	}

	for i := 1; i <= nb; i++ {
		a = a * i
	}
	return a
}
