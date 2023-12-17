package piscine

import "fmt"

func PrintNum(num int) {
	n := '0'
	for a := 1; a <= num%10; a++ {
		n++
	}
	for a := -1; a >= num%10; a-- {
		n++
	}
	if num%10 != 0 {
		PrintNum(num % 10)
	}
	fmt.Print(n)
}

func PrintNbr(nbr int) {
	if nbr < 0 {
		fmt.Print('-')
	}
	PrintNum(nbr)
}
