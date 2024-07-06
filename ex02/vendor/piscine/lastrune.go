package piscine

func Strlen(s string) int {
	var i int = 0
	for range s {
		i++
	}
	return i
}

func LastRune(s string) rune {
	sLen := Strlen(s)
	if sLen == 0 {
		return 0
	}
	return rune(s[sLen-1])
}
