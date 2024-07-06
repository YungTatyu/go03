package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Join([]string{"Hello!", " How", " are", " you?"}, "sep"))
	fmt.Println(piscine.Join([]string{"hey!", "", " hey", " star"}, ":"))
	fmt.Println(piscine.Join([]string{"", "", "first"}, ""))
	fmt.Println(piscine.Join([]string{"", "", ""}, ""))
	fmt.Println(piscine.Join([]string{"", "", ""}, "t"))
}
