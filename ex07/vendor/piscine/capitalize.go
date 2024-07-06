package piscine

const (
	swStart = iota
	swNonWord
	swWord
)

const diff = 'a' - 'A'

func IsAlpha(r rune) bool {
	return IsUpperCase(r) || IsLowerCase(r)
}

func IsDigid(r rune) bool {
	return r >= '0' && r <= '9'
}

func IsAlphaNum(r rune) bool {
	return IsAlpha(r) || IsDigid(r)

}

func Capitalize(s string) string {
	var re string
	state := swNonWord
	for _, r := range s {
		switch state {
		case swStart:
			state = swNonWord
		case swNonWord:
			if IsAlphaNum(r) {
				state = swWord
			}
			re += string(ToUpper(string(r)))
		case swWord:
			if !IsAlphaNum(r) {
				state = swNonWord
			}
			re += string(ToLower(string(r)))
		}
	}
	return re
}
