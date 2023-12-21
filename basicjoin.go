package piscine

func BasicJoin(elems []string) string {
	var result string
	for _, words := range elems {
		result += words
	}
	return result
}

// func main() {
// 	elems := []string{"Hello!", " How", " are", " you?"}
// 	fmt.Println(BasicJoin(elems))
// }
