package module01

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//	BaseToDec("E", 16) => 14
//	BaseToDec("1110", 2) => 14
//
// base 2 = max digit is 1
// base 3 = max digit is 2
// ...
// base 10 = max digit is 9
// ...
// base 16 = max digit is F (15)
// 10 = A
// 11 = B
// 12 = C
// 13 = D
// 14 = E
// 15 = F
//
// 10011 base 2 = 1*2^4 + 0*2^3 + 0*2^2 + 1*2^1 + 1*2^0 = 16 + 0 + 0 + 2 + 1 = 19 base 10
//method 1
//func BaseToDec(value string, base int) int {
// 	result := 0
// 	for i := 0; i < len(value); i++ {
// 		ch := value[len(value)-1-i]
// 		if ch >= '0' && ch <= '9' {
// 			v := int(ch - '0')
// 			result += v * int(math.Pow(float64(base), float64(i)))
// 		} else if 'A' <= ch && ch <= 'F' {
// 			a := int(ch - 'A' + 10)
// 			result += a * int(math.Pow(float64(base), float64(i)))
// 		}
// 	}
// 	return result
// }

// method2 101
// func BaseToDec(value string, base int) int {
// 	//const charset = "0123456789ABCDEF"
// 	result := 0
// 	multiplier := 1
// 	for i := len(value) - 1; i >= 0; i-- {
// 		ch := value[i]
// 		var digit int
// 		if ch >= '0' && ch <= '9' {
// 			digit = int(ch - '0')
// 		} else if 'A' <= ch && ch <= 'F' {
// 			digit = int(ch - 'A' + 10)
// 		} else {
// 			panic("invalid")
// 		}
// 		result += digit * multiplier
// 		multiplier = multiplier * base
// 	}
// 	return result
// }

// method3 100
func BaseToDec(value string, base int) int {
	const charset = "0123456789ABCDEF"
	result := 0
	multiplier := 1
	for i := len(value) - 1; i >= 0; i-- {
		index := -1
		for j, char := range charset {
			if char == rune(value[i]) {
				index = j
				break
			}
		}
		if index < 0 {
			panic("something went wrong")
		}
		result += index * multiplier
		multiplier = multiplier * base
	}
	return result
}
