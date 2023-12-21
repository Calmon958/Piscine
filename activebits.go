package piscine

func ActiveBits(n int) int {
	bits := 0
	for n != 0 {
		if n&1 == 1 {
			bits++
		}
		n = n >> 1
	}
	return bits
}
