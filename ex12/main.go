package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsPrintable("Hello! How are you?"))
	fmt.Println(piscine.IsPrintable("Gopher[]09$^*&"))
	fmt.Println(piscine.IsPrintable("helllo"))
	fmt.Println(piscine.IsPrintable("heeloo\n"))
	fmt.Println(piscine.IsPrintable("\t"))
	fmt.Println(piscine.IsPrintable(""))
}
