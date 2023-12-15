package piscine

import "fmt"

func main() {
	sentence := "Hello world"
	for index, letter := range sentence {
		fmt.Printf("The index is %v letter %c", index, letter)
	}
}
