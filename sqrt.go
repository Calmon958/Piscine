package piscine

func Sqrt(nb int) int {
	sqrt := 0
	for i := 0; i*i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return sqrt
}
