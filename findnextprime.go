package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	for i := nb; ; i++ {
		if IsPrime2(i) {
			return i
		}
	}
}

func IsPrime2(nb int) bool {
	if nb < 2 {
		return false
	} else if nb == 2 {
		return true
	} else if nb%2 == 0 && nb > 2 {
		return false
	}
	for i := 3; i*i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(FindNextPrime(97197))
// }
