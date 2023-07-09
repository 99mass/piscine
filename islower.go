package piscine

func IsLower(s string) bool {
	IsLower := 0
	for i := 0; i <= len(s)-1; i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			IsLower = IsLower + 1
		}
	}
	if IsLower == len(s) {
		return true
	} else {
		return false
	}
}
