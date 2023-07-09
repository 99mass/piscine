package piscine

func ActiveBits(n int) int {
	if n >= 0 && n < 2 {
		return n
	}
	return (n % 2) + ActiveBits(n/2)
}
