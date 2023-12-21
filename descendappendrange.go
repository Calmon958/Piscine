package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return []int{}
	}
	var value []int
	for x := max; x > min; x-- {
		value = append(value, x)
	}
	return value
}
