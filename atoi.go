package piscine

func Atoi(s string) int {
	r := 0
	sign := 1
	index := 0
	if s[0] == '-' {
		sign = -1
		index = 1
	}
	for i := index; i < len(s); i++ {
		a := int(s[i] - '0')
		if a >= 0 && a <= 9 {
			r *= 10
			r += a
		} else {
			return 0
		}
	}
	return r * sign
}
