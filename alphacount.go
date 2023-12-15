package piscine

func AlphaCount(s string) int {
	count := 0
	for _, word := range s {
		if (word >= 'a' && word <= 'z') || (word >= 'A' && word <= 'Z') {
			count++
		}
	}
	return count
}
