package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	table(args)
	for _, val := range args {
		for _, val2 := range val {
			z01.PrintRune(rune(val2))
		}
		z01.PrintRune('\n')
	}
}

func table(tab []string) { // sorting through slice
	for a := 0; a < len(tab)-1; a++ {
		for b := a + 1; b < len(tab); b++ {
			if tab[a] > tab[b] {
				tab[a], tab[b] = tab[b], tab[a]
			}
		}
	}
}
