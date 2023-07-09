package piscine

func Map(f func(int) bool, a []int) []bool {
	boolean := []bool{}
	for _, z := range a {
		boolean = append(boolean, f(z))
	}
	return boolean
}
