package piscine

func Alphabet(l rune) bool {
	if (l >= 'A' && l <= 'Z') || (l >= 'a' && l <= 'z') || (l >= '0' && l <= '9') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	a := []rune(s)

	r := true

	for i := 0; i < len(s); i++ {
		if Alphabet(a[i]) == true && r {
			if a[i] >= 'a' && a[i] <= 'z' {
				a[i] = 'A' - 'a' + a[i]
			}
			r = false
		} else if a[i] >= 'A' && a[i] <= 'Z' {
			a[i] = 'a' - 'A' + a[i]
		} else if Alphabet(a[i]) == false {
			r = true
		}
	}
	return string(a)
}
