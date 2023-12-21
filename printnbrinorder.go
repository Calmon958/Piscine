package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	count := make([]int, 10) // Array to count occurrences of each digit from 0 to 9

	if n == 0 {
		z01.PrintRune('0') // Print '0' for negative input
		return
	}

	// Count occurrences of each digit in the number
	for n > 0 {
		num := n % 10
		count[num]++
		n /= 10
	}

	// Print digits in ascending order
	for i := 0; i < 10; i++ {
		for j := 0; j < count[i]; j++ {
			srune := '0' + rune(i)

			z01.PrintRune(srune) // Print the digit by adding '0' to the digit value
		}
	}
	// z01.PrintRune('\n')
}

// func main() {
// 	PrintNbrInOrder(321)
// 	PrintNbrInOrder(0)
// 	PrintNbrInOrder(321)
// }
