package piscine

func parseBase(base string) bool {
	if Strlen(base) < 2 {
		return false
	}
	for i, v := range base {
		if v == '+' || v == '-' {
			return false
		}
		if Index(base[i+1:], string(v)) != -1 {
			return false
		}
	}
	return true
}

func AtoiBase(s string, base string) int {
	if !parseBase(base) {
		return 0
	}
	var num int = 0
	len := int(Strlen(base))
	for _, v := range s {
		index := Index(base, string(v))
		num = num*len + index
	}
	return num
}
