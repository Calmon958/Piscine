package main

import "fmt"

func main() {
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				for d := 0; d <= 9; d++ {
					if a > c || (a == c && b >= d) {
						continue
					}
					fmt.Print(a)
					fmt.Print(b)
					fmt.Print(" ")
					fmt.Print(c)
					fmt.Print(d)

					if !(a == 9 && b == 9 && c == 8 && d == 9) {
						fmt.Print(",")
						fmt.Print(" ")
					}
				}
			}
		}
	}
	fmt.Println()
}
