package main

import (
	"ft"
	"piscine"
)

const (
	UintMax = ^uint(0)
	UintMin = 0
	IntMax  = int(UintMax >> 1)
	IntMin  = -IntMax - 1
)

func main() {
	piscine.PrintNbrBase(125, "0123456789")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(-125, "01")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(125, "0123456789ABCDEF")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(-125, "choumi")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(125, "aa")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(IntMax, "0123456789ABCDEF")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(IntMin, "0123456789ABCDEF")
	ft.PrintRune('\n')
}
