package piscine

func ShoppingListSort(slice []string) []string {
	// slice := []string{}
	for a := 0; a < len(slice)-1; a++ {
		for b := a + 1; b < len(slice); b++ {
			if len(slice[a]) > len(slice[b]) {
				slice[a], slice[b] = slice[b], slice[a]
			}
		}
	}
	return slice
}
