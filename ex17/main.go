package main

import (
	"fmt"
	"piscine"
)

const (
	UintMax = ^uint(0)
	UintMin = 0
	IntMax  = int(UintMax >> 1)
	IntMin  = -IntMax - 1
)

func main() {
	fmt.Println(piscine.AtoiBase("125", "0123456789"))
	fmt.Println(piscine.AtoiBase("1111101", "01"))
	fmt.Println(piscine.AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(piscine.AtoiBase("uoi", "choumi"))
	fmt.Println(piscine.AtoiBase("bbbbbab", "-ab"))
}
