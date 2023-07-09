package piscine

func Index(s string, toFind string) int {
	taille := 0
	tailleSup := 0

	for i := range s {
		taille = i + 1
	}

	for i := range toFind {
		tailleSup = i + 1
	}

	t := 0

	for i := 0; i < taille; i++ {
		j := 0
		t = i
		for j < tailleSup {
			if t < taille && s[t] == toFind[j] {
				j++
				t++
			} else {
				break
			}
		}
		if j == tailleSup {
			return i
		}
	}
	return -1
}
