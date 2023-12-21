package piscine

import (
	"strings"

	"github.com/01-edu/z01"
)

func valBase(b string) bool {
	if len(b) < 2 {
		return false
	}
	specialChar := make(map[rune]bool)
	for _, val := range b {
		if val == '+' || val == '-' {
			return false
		}
		if specialChar[val] {
			return false
		}
		specialChar[val] = true
	}
	return true
}

func gIndex(nb byte, baze string) int {
	for x, index := range baze {
		if byte(index) == nb {
			return x
		}
	}
	return -1
}

func PrintNbrBase(nbr int, base string) {
	if !valBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		// z01.PrintRune('\n')
		return
	}
	neg := nbr < 0
	if neg {
		nbr = -nbr
	}
	var result strings.Builder
	for nbr > 0 {
		remind := nbr % len(base)
		result.WriteByte(base[remind])
		nbr /= len(base)
	}
	if neg {
		result.WriteByte('-')
	}

	finString := result.String()
	srune := []rune(finString)
	for a, b := 0, len(srune)-1; a < b; a, b = a+1, b-1 {
		srune[a], srune[b] = srune[b], srune[a]
	}
	for _, r := range srune {
		z01.PrintRune(r)
	}
	// z01.PrintRune('\n')
}

// func main() {
// 	PrintNbrBase(125, "0123456789")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(-125, "01")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(125, "0123456789ABCDEF")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(-125, "choumi")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(125, "aa")
// 	z01.PrintRune('\n')
// }
