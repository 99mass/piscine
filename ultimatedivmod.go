package piscine

func UltimateDivMod(a *int, b *int) {
	y := *a / *b
	z := *a % *b
	*a = y
	*b = z
}
