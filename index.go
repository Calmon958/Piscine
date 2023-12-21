package piscine

func Index(s string, toFind string) int {
	// lenFind := len(toFind)
	for a := 0; a <= len(s)-len(toFind); a++ {
		match := true
		for b := 0; b < len(toFind); b++ {
			if s[a+b] != toFind[b] {
				match = false
				break
			}
		}
		if match {
			return a
		}
	}
	return -1
}

// func main() {
// 	fmt.Println(Index("Hello!", "l"))
// 	fmt.Println(Index("Salut!", "alu"))
// 	fmt.Println(Index("Ola!", "hOl"))
// }
