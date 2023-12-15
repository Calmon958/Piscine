package piscine

func Unmatch(x []int) int {
	for a := 0; a < len(x)-1; a++ {
		for b := a + 1; b < len(x); b++ {
			if x[a] > x[b] {
				x[a], x[b] = x[b], x[a]
			}
		}
	}
	for len(x) > 0 {
		if len(x)-1 != 0 {
			if x[0] == x[1] {
				x = x[2:]
			} else {
				return x[0]
			}
		} else {
			return x[0]
		}
	}
	// if x[a] != x[b] && nbrs[a]+nbrs[b] {
	// 	return nbrs[a], nbrs[b]
	// } else {
	// 	continue
	// }
	return -1
}
