package piscine

func IsPrintable(s string) bool {
	for _, ran := range s {
		if ran >= ' ' && ran <= '~' {
			continue
		} else {
			return false
		}
	}
	return true
}
