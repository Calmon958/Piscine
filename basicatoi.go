package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, i := range s { // index value for the digit
		x := int(i - '0') // converts rune to int
		result = result*10 + x
	}
	return result
}
