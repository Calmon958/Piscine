package piscine

func NRune(sentence string, n int) rune {
	if n <= 0 || n > len(sentence) {
		return 0
	}
	result := []rune(sentence)
	return result[n-1]
}

// func main() {
// 	NRune("Hello", 4)
// }
