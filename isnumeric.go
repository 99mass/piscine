package piscine

func GetState(value rune) bool {
	if value < '0' || value > '9' {
		return false
	}
	return true
}

func IsNumeric(s string) bool {
	chaine := []rune(s)
	for _, v := range chaine {
		if GetState(rune(v)) == false {
			return false
		}
	}
	return true
}
