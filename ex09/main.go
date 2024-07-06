package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsNumeric("Hello! How are you?"))
	fmt.Println(piscine.IsNumeric("HelloHowareyou"))
	fmt.Println(piscine.IsNumeric("What’s this 4?"))
	fmt.Println(piscine.IsNumeric("Whatsthis4"))
	fmt.Println(piscine.IsNumeric("0083031373"))
	fmt.Println(piscine.IsNumeric(""))
}
