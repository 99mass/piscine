package piscine

func ToLower(s string) string {
	lowerChaine := []rune(s)
	for i, charr := range lowerChaine {
		if charr >= 'A' && charr <= 'Z' {
			lowerChaine[i] = charr + ('a' - 'A')
		}
	}
	return string(lowerChaine)
}
