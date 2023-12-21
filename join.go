package piscine

func Join(strs []string, sep string) string {
	var answer string
	for a, val := range strs {
		answer += val
		if a != len(strs)-1 {
			answer += sep
		}
	}
	return answer
}
