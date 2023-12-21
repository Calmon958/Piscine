package piscine

func IsUpper(s string) bool {
	// var result bool
	for _, val := range s {
		if val >= 'A' && val <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsUpper("HELLO"))
// 	fmt.Println(IsUpper("HELLO!"))
// 	fmt.Println(IsUpper("Hello!"))
// 	fmt.Println(IsUpper("+s8ZH7zX*6+9R"))
// }
