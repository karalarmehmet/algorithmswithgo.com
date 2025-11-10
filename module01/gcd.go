package module01

// A=30 ,B=15
// 30=[2,3,5]
// 15=[3,5]
// GCD=[3,5]

// A=100 ,B=8
// 100=[2,2,5,5]
// 8=[2,2,2]
// GCD=[2,2]

//	func GCD(a, b int) int {
//		Primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
//		res := Factor(Primes, a)
//		res2 := Factor(Primes, b)
//		var common []int
//		var product int
//		product = 1
//		for i := 0; i < len(res); i++ {
//			for j := 0; j < len(res2); j++ {
//				if res[i] == res[j] {
//					common = append(common, res[i])
//				}
//			}
//		}
//		for _, n := range common {
//			product *= n
//		}
//		return product
//	}
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
