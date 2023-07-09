package piscine

func StrLen(s string) int {
	compteur := 0
	theRune := []rune(s)
	for index := range theRune {
		compteur = index + 1
	}
	return compteur
}
