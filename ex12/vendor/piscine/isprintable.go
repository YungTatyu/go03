package piscine

func Printable(r rune) bool {
	return r >= 32 && r <= 126
}

func IsPrintable(s string) bool {
	for _, v := range s {
		if !Printable(v) {
			return false
		}
	}
	return true
}
