package piscine

func Rot14(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'a' && runes[i] <= 'l' || runes[i] >= 'A' && runes[i] <= 'L' {
			runes[i] = runes[i] + 14
		} else if runes[i] >= 'm' && runes[i] <= 'z' || runes[i] >= 'M' && runes[i] <= 'Z' {
			runes[i] = runes[i] - 12
		}
	}
	return string(runes)
}
