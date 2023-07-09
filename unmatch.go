package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				a[i] = -1
				a[j] = -1
				break
			}
		}
	}
	for k := 0; k < len(a); k++ {
		if a[k] != -1 {
			return a[k]
		}
	}
	return -1
}
