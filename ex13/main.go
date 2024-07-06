package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Concat("Hello!", "How are you?"))
	fmt.Println(piscine.Concat("", ""))
	fmt.Println(piscine.Concat("hi", ""))
	fmt.Println(piscine.Concat("", "hi"))
}
