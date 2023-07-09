package piscine

func IsPrintable(s string) bool {
	chaine := []rune(s)

	for i := 0; i <= len(s)-1; i++ {
		if chaine[i] >= 0 && chaine[i] <= 31 {
			return false
		}
	}
	return true
}
