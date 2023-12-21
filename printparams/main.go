package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for _, nb := range args {
		for _, val := range nb {
			z01.PrintRune(val)
		}
		z01.PrintRune('\n')
	}
}
