package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintString(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func PrintNumber(n int) {
	if n < 0 || n > 9 {
		return
	}
	slice := []rune{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	x := slice[n]
	z01.PrintRune(x)
}

func PrintVal(nb int) {
	if nb < 0 {
		z01.PrintRune('-')
		nb = -nb
	}
	if nb < 10 {
		PrintNumber(nb)
	} else {
		PrintVal(nb / 10)
		PrintVal(nb % 10)
	}
}

func main() {
	points := &point{}

	setPoint(points)
	x, y := points.x, points.y
	PrintString("x = ")
	PrintVal(x)
	PrintString(", y = ")
	PrintVal(y)
	z01.PrintRune('\n')
}
