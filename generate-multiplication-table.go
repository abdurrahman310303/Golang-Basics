package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", n, i, n*i)
	}
	fmt.Println("Multiplication table generated successfully")
}
