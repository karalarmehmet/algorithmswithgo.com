package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {

	// var res string
	// for i:=0 ; i<len(word);i++{
	// 	res= string(word[i]) + res
	// }
	// return res

	// var sb strings.Builder
	// for i:= len(word)-1 ; i>=0 ; i--{
	// 	sb.WriteByte(word[i])
	// }
	// return sb.String()

	// var res string
	// for i:= len(word)-1 ; i>=0 ; i--{
	// 	res = res + string(word[i])
	// }
	// return res

	var res string
	for _, val := range word {
		res = string(val) + res
	}
	return res

}
