package piscine

import "fmt"

func main() {
	for a := 0; a <= 9; a++ {
		for b := 1; b <= 8; b++ {
			for c := 2; c <= 7; c++ {
				fmt.Print(a)
				fmt.Print(b)
				fmt.Print(c)
				if a == 9 && b == 8 && c == 7 {
					fmt.Print(".")
				}
				if !(a == 9 && b == 8 && c == 7) {
					fmt.Print(",")
					fmt.Print(" ")
				}
			}
		}
	}
}
