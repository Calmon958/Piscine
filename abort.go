package piscine

func Abort(a, b, c, d, e int) int {
	//	int x
	int_x := []int{a, b, c, d, e}
	for a := 0; a < len(int_x)-1; a++ {
		for b := a + 1; b < len(int_x); b++ {
			if int_x[a] > int_x[b] {
				int_x[a], int_x[b] = int_x[b], int_x[a]
			}
		}
	}
	return int_x[2]
}

// 	digits := []int{a, b, c, d, e}
// 	sum := digits[0]
// 	min := digits[0]
// 	max := digits[0]
// 	for i := 0; i < len(digits); i++ {
// 		if digits[i] < min[i] {
// 			return min[i]
// 		} else if digits[i] > max[i] {
// 			return max[i]
// 		} else {
// 			sum += digits[i]
// 		}
// 	}
// 	median := sum - min - max
// 	return median
// }
