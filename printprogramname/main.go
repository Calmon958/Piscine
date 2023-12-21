package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[0]
	// if len(args) > 0 {
	// nb := args[0]
	for _, i := range args[2:] {
		z01.PrintRune(i)
	}
	//}
	z01.PrintRune('\n')
}
