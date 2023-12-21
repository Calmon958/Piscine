package piscine

import "github.com/01-edu/z01"

func PrintNum(nb int) {
	a := '0'
	// if nb == 0 {
	// 	z01.PrintRune('0')
	// }
	for b := 1; b <= nb%10; b++ {
		a++
	}
	for b := -1; b >= nb%10; b-- {
		a++
	}
	if nb/10 != 0 {
		PrintNum(nb / 10)
	}
	z01.PrintRune(a)
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	PrintNum(n)
}

// 	num := rune(n)
// 	if num < '0' {
// 		z01.PrintRune('-')
// 		z01.PrintRune(rune(num))
// 	} else {
// 		z01.PrintRune(rune(num))
// 	}
// }

// func main() {
// 	PrintNbr(-123)
// 	PrintNbr(0)
// 	PrintNbr(123)
// 	z01.PrintRune('\n')
// }
