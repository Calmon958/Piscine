package piscine

func LastRune(sentence string) rune {
	result := []rune(sentence)
	return result[len(result)-1]
}
