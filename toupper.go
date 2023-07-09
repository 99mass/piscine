package piscine

func ToUpper(s string) string {
	upperChaine := []rune(s)
	for i, charr := range upperChaine {
		if charr >= 'a' && charr <= 'z' {
			upperChaine[i] = charr - ('a' - 'A')
		}
	}
	return string(upperChaine)
}
