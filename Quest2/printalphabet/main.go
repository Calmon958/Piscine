package piscine

import "fmt"

func main() {
	// i := 'a'
	// for i := range s{
	for i := 'a'; i <= 'z'; i++ {
		fmt.Print(i)
		fmt.Print(",")
	}
	fmt.Println()
}
