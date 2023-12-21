package piscine

func Max(a []int) int {
	max := 0
	if len(a) > 0 {
		max = a[0]
		for _, b := range a {
			if b > max {
				max = b
			}
		}
	}
	return max
}
