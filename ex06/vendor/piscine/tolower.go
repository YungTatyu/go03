package piscine

const diff = 'a' - 'A'

func IsUpperCase(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func ToLower(s string) string {
	var i int = 0
	var re string
	for range s {
		r := rune(s[i])
		i++
		if IsUpperCase(r) {
			r += diff
		}
		re += string(r)
	}
	return re
}
