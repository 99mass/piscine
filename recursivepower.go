package piscine

func RecursivePower(nb int, power int) int {
	a := 1
	if power < 0 || power > 20 {
		return 0
	} else if power == 0 {
		return 1
	}
	a = nb * RecursivePower(nb, (power-1))

	return a
}
