package piscine

func ReverseMenuIndex(menu []string) []string {
	reversed := make([]string, len(menu))
	for i := range menu {
		reversed[len(menu)-i-1] = menu[i]
	}
	return reversed
}
