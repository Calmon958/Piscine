package piscine

func DealAPackOfCards(deck []int) [4][3]int {
	cards := len(deck) / 4
	var result [4][3]int

	for i := 0; i < 4; i++ {
		player := i + 1
		result[i][0] = player

		for j := 0; j < cards; j++ {
			index := i*cards + j
			result[i][j+1] = deck[index]
		}
	}

	return result
}
