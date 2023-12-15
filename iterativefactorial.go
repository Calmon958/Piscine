package piscine

func IterativeFactorial(nb int) int {
	factorial := 1
	if nb == 1 || nb == 0 {
		return 1 // the integer number is 0 or 1 it returns 1
	} else if nb > 1 && nb <= 30 { // the integer is between the higher limit and the lower limit
		for a := 1; a <= nb; a++ {
			factorial = factorial * a
		}
	} else {
		return 0 // if the integer beyond the limits
	}
	return factorial
}
