package piscine

func IsLowerCase(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func IsLower(s string) bool {
	for _, v := range s {
		if !IsLowerCase(v) {
			return false
		}
	}
	return true
}
