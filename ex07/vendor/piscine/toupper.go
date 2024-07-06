package piscine

func IsLowerCase(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func ToUpper(s string) string {
	var i int = 0
	var re string
	for range s {
		r := rune(s[i])
		i++
		if IsLowerCase(r) {
			r -= diff
		}
		re += string(r)
	}
	return re
}
