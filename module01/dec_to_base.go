package module01

import "strings"

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
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

// func DecToBase(dec, base int) string {
// 	const digits = "0123456789ABCDEF"
// 	var res string
// 	for dec > 0 {
// 		rem := dec % base
// 		res = string(digits[rem]) + res
// 		dec = dec / base
// 	}
// 	return res
// }

func DecToBase(dec, base int) string {
	const digits = "0123456789ABCDEF"
	sb := strings.Builder{}

	for dec > 0 {
		rem := dec % base
		sb.WriteByte(digits[rem])
		dec = dec / base
	}

	result := sb.String()
	runes := []rune(result)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
