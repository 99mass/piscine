package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	} else {
		k := 0
		for i := 0; i < len(a); i++ {
			k = i
			for j := i + 1; j < len(a); j++ {
				if a[j] > a[k] {
					a[k] = a[j]
				}
			}
			a[k], a[i] = a[i], a[k]
		}
		return a[0]
	}
}
