package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {

	sum := 0
	for _, val := range numbers {
		sum = sum + val
	}
	return sum

	// fmt.Println(numbers)
	// if len(numbers) == 0{
	// 	return 0
	// }
	// return numbers[0] + Sum(numbers[1:])

	// sum(3,4,5)
	// 3 + sum(4,5)
	//      4 + sum(5)
	// 		5 + sum()
	// 		     0
}
