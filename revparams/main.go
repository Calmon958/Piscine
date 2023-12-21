package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	point := len(args) - 1
	for point >= 0 {
		for _, nb := range args[point] {
			z01.PrintRune(nb)
		}
		point--
		z01.PrintRune('\n')
	}
}
