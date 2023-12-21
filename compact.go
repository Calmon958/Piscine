package piscine

func Compact(ptr *[]string) int {
	var list []string
	for _, v := range *ptr {
		if v != "" {
			list = append(list, v)
		}
	}
	*ptr = list
	return len(list)
}
