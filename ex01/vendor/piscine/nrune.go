package piscine

func Strlen(s string) int {
	var i int = 0
	for range s {
		i++
	}
	return i
}

func NRune(s string, n int) rune {
	if n <= 0 || n > Strlen(s) {
		return 0
	}
	return rune(s[n-1])
}
