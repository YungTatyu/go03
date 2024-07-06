package piscine

func Alpha(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func Digid(r rune) bool {
	return r >= '0' && r <= '9'
}

func IsAlphaNum(r rune) bool {
	return Alpha(r) || Digid(r)

}

func IsAlpha(s string) bool {
	for _, v := range s {
		if !IsAlphaNum(v) {
			return false
		}
	}
	return true
}
