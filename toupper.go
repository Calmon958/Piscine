package piscine

func ToUpper(s string) string {
	br := []rune(s)
	for x, val := range br {
		if val >= 'a' && val <= 'z' {
			br[x] = val - ('a' - 'A')
		}
	}
	return string(br)
}

// func main() {
// 	fmt.Println(ToUpper("Hello! How are you?"))
// 	fmt.Println(ToUpper("Yooh Nigga whats up"))
// }
