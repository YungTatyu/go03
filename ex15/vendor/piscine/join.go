package piscine

func Strlen(s string) uint {
	var i uint = 0
	for range s {
		i++
	}
	return i
}

func Join(strs []string, sep string) string {
	var re string
	for _, v := range strs {
		re += (v + sep)
	}
	return re[:Strlen(re)-Strlen(sep)]
}
