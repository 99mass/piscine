package piscine

func Sqrt(nb int) int {
	if nb > 0 {
		if nb == 0 {
			return 0
		}
		for i := 0; i*i <= nb; i++ {
			if i*i == nb {
				return i
			}
		}
	} else {
		return 0
	}
	return 0
}
