package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.BasicJoin([]string{"Hello!", " How", " are", " you?"}))
	fmt.Println(piscine.BasicJoin([]string{"hey!", "", " hey", " star"}))
	fmt.Println(piscine.BasicJoin([]string{"", "", "first"}))
}
