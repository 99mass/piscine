package piscine

func TrimAtoi(s string) int {
	var a []int
	r := 0
	min := -1
	b := 0
	i := 0
	c := 0
	for _, rune := range s {
		if rune == '-' {
			min = i
		}
		if Myfonction(rune) {
			if b == 0 {
				b = i
			}
			a = append(a, int(rune-'0'))
		}
		i++
	}

	for count := range a {
		c = count + 1
	}

	for i := 0; i < c; i++ {
		r = r*10 + a[i]
	}

	if min < b && min != -1 {
		r = r * -1
	}
	return r
}

func Myfonction(x rune) bool {
	for a := '0'; a <= '9'; a++ {
		if x == a {
			return true
		}
	}
	return false
}
