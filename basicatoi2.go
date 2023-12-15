package piscine

func BasicAtoi2(s string) int {
	ans := 0
	for _, i := range s {
		x := int(i - '0') // used to convert rune to integer value
		if x < 0 || x > 9 {
			return 0 // returns zero when character isn't of type int
		}
		ans = ans*10 + x
	}
	return ans
}
