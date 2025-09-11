package module01

import "fmt"
// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	// for i := 1 ; i < n ; i++ {
	// 	printFizzBuzzValue(i)
	// 	fmt.Print(", ")
	// }
    // printFizzBuzzValue(n)
	// fmt.Println()
for i := 1 ; i <= n ; i++ {
	switch{
		case i % 3 == 0 && i % 5 == 0: // it means i % 15 == 0 
			fmt.Print("Fizz Buzz")
		case i % 3 == 0:  
			fmt.Print("Fizz")	
		case i % 5 == 0:
			fmt.Print("Buzz")
		default:
			fmt.Print(i)	
		}
		if i!=n{
			fmt.Print(", ")
		}	
	}
	fmt.Println()
}	

// func printFizzBuzzValue(n int ){
// 	// if n % 3 == 0 && n % 5 == 0{
// 	// 	fmt.Print("Fizz Buzz")
// 	// }else if n % 3 == 0 {
// 	// 	fmt.Print("Fizz")
// 	// }else if n % 5 == 0 {
// 	// 	fmt.Print("Buzz")	
// 	// } else {
// 	// 	fmt.Print(n)
// 	// }
// 	switch{
// 		case n % 3 == 0 && n % 5 == 0: // it means i % 15 == 0 
// 			fmt.Print("Fizz Buzz")
// 		case n % 3 == 0:  
// 			fmt.Print("Fizz")	
// 		case n % 5 == 0:
// 			fmt.Print("Buzz")
// 		default:
// 			fmt.Print(n)	
// 		}
// }