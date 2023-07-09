package piscine

func LastRune(s string) rune {
	lastChar := []rune(s)
	return lastChar[len(lastChar)-1]
}
