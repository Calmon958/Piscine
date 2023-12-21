package piscine

func ToLower(s string) string {
	br := []rune(s)
	for a, index := range br {
		if index >= 'A' && index <= 'Z' {
			br[a] = index + ('a' - 'A')
		}
	}
	return string(br)
}
