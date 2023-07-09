package piscine

func IterativePower(nb int, power int) int {
	a := 1
	if power < 0 || power > 20 {
		return 0
	} else if power == 0 {
		return 1
	}
	for i := 1; i <= power; i++ {
		a = a * nb
	}

	return a
}
