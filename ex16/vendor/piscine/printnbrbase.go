package piscine

import "ft"

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

func recurPrintNbrBase(nbr uint, base string) {
	if nbr == 0 {
		return
	}
	len := Strlen(base)
	recurPrintNbrBase(nbr/len, base)
	ft.PrintRune(rune(base[nbr%len]))
}

func PrintNbrBase(nbr int, base string) {
	if !parseBase(base) {
		ft.PrintRune('N')
		ft.PrintRune('V')
		return
	}
	var num uint = uint(nbr)
	if nbr < 0 {
		ft.PrintRune('-')
		num = uint(-nbr)
	}
	recurPrintNbrBase(num, base)
}
