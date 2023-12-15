package piscine

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
	/*	for index, letter := range []byte(a, b)  {
				if letter != b[index] {
					if letter < b[index] {
						return -1
					} else {
						return 1
					}
				}
			}
			return 0
		}

		if a == b {
			return 0
		} else if a
		}
	*/
}

// func main() {
// 	Compare("Hello", "World")
// }
