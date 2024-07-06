package piscine

func Digid(r rune) bool {
	return r >= '0' && r <= '9'
}

func IsNumeric(s string) bool {
	for _, v := range s {
		if !Digid(v) {
			return false
		}
	}
	return true
}
