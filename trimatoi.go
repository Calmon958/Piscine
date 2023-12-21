package piscine

func TrimAtoi(input string) int {
	const (
		maxInt = 1<<31 - 1
		minInt = -1 << 31
	)

	var num int
	var neg bool
	for _, value := range input {
		if value >= '0' && value <= '9' {
			digit := int(value - '0')
			num = num*10 + digit
			// num += string(value)
		} else if value == '-' && num == 0 {
			neg = true
		}
	}
	if neg {
		num = -num
	}
	if num > maxInt {
		return maxInt
	} else if num < minInt {
		return minInt
	}
	return num
}

/*if num == "" {
		return 0
	}
	result, err := strconv.Atoi(num)
	if err != nil {
		return 0
	}
	if neg {
		return -result
	}
	return result
}*/

// func main() {
// 	fmt.Println(TrimAtoi("12345"))
// 	fmt.Println(TrimAtoi("str123ing45"))
// 	fmt.Println(TrimAtoi("012 345"))
// 	fmt.Println(TrimAtoi("Hello World!"))
// 	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
// 	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
// 	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
// 	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
// }
