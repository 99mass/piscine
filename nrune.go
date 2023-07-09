package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}
	runes := []rune(s)
	taille := 0
	for i := range s {
		taille = i + 1
	}
	if taille == 0 || n > taille {
		return 0
	}

	return runes[n-1]
}
