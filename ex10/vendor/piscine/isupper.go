package piscine

const diff = 'a' - 'A'

func IsUpperCase(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func IsUpper(s string) bool {
	for _, v := range s {
		if !IsUpperCase(v) {
			return false
		}
	}
	return true
}
