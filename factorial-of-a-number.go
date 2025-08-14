package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Println(factorial)
}
