package main

import (
	"algo/module01"
	"fmt"
)

func main() {
	var n int
	fmt.Println("How many inputs ? ")
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		gcd := module01.GCD(a, b)
		fmt.Println(gcd)
	}
}
