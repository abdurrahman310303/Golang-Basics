package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println(sum)
}
