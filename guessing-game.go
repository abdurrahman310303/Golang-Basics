package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.NewSource(time.Now().UnixNano())

	secretNumber := rand.Intn(100) + 1

	fmt.Println("welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100")
	fmt.Println("Can you guess it?")

	attempts := 0
	maxAttempts := 7

	for attempts < maxAttempts {
		attempts++
		fmt.Printf("\nAttempt %d/%d: ", attempts, maxAttempts)

		var guess int
		fmt.Scanln(&guess)

		if guess == secretNumber {

			fmt.Printf("Congratulations! you guessed it in %d attempts!\n", attempts)
			return
		} else if guess < secretNumber {
			fmt.Println("Too low! try a higher Number")
		} else {
			fmt.Println("Too high! try a lower Number")
		}
	}

	fmt.Println("Game Over! The number was %d.\n", secretNumber)
}
