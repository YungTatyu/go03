package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsLower("Hello! How are you?"))
	fmt.Println(piscine.IsLower("Gopher[]09$^*&"))
	fmt.Println(piscine.IsLower("helllo"))
	fmt.Println(piscine.IsLower("heeloo!"))
	fmt.Println(piscine.IsLower(""))
}
