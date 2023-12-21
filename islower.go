package piscine

func IsLower(s string) bool {
	for _, val := range s {
		if val >= 'a' && val <= 'z' {
			continue
		} else {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsLower("hello"))
// 	fmt.Println(IsLower("hello!"))
// }
