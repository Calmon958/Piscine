package piscine

func Capitalize(s string) string {
	Upper := func(s byte) byte {
		if s >= 'a' && s <= 'z' {
			return s - ('a' - 'A')
		}
		return s
	}

	Lower := func(s byte) byte {
		if s >= 'A' && s <= 'Z' {
			return s + ('a' - 'A')
		}
		return s
	}

	Alpha := func(index byte) bool {
		return (index >= 'a' && index <= 'z') || (index >= 'A' && index <= 'Z') || (index >= '0' && index <= '9')
	}
	runes := []byte(s)
	for a := range runes {
		if Alpha(runes[a]) {
			if a == 0 || !Alpha(runes[a-1]) {
				runes[a] = Upper(runes[a])
			} else {
				runes[a] = Lower(runes[a])
			}
		}
	}
	return string(runes)
}

// s = Lower(s)
// for a, value := range s {
// 	if a == 0 {
// 		s = Upper(string(value)) + s[a+1:]
// 	} else {
// 		if Alpha(string(value)) && !Alpha(string(s[a-1])) {
// 			if a != len(s)-1 {
// 				s = s[:a] + Upper(string(value)) + s[a+1:]
// 			} else {
// 				s = s[:1] + Upper(string(value))
// 			}
// 			}
// 		}
// 	}
// 	return s
// }

// func main() {
// 	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
// 	fmt.Println(Capitalize("7<r=T7iI~#r@C"))
// }
