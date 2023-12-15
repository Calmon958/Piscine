package piscine

func FirstRune(sentence string) rune {
	result := []rune(sentence)
	return result[0]
}

/*	for _, letter := range []rune(sentence) {
		fmt.Printf("%c\n", letter)
	}
	return rune(sentence[0])
}
*/
/*func FirstRune(letter string) {
letter := "Hello"
var index int
sentence := []byte("Hello")
counter := 0
for _, letter := range letter {
	counter++
	fmt.Printf("index: %v  letter: %c\n", index, letter)
}
fmt.Printf("Counter value: %v\n", counter)
fmt.Println(sentence)
return letter*/
// func main() {
// 	//	s := "Hello"
// 	fmt.Println(FirstRune("Hello"))
// 	//	fmt.Println(s)
// }
