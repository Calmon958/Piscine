package main

import "fmt"

func main() {
	for i := 'z'; i >= 'a'; i-- {
		fmt.Print(i)
		fmt.Print(",")
	}
	fmt.Println()
}
