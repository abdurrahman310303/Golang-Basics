package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Switch Selection Structure")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You can move 1 spot")
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
		fallthrough
	case 4:
		fmt.Println("You can move 4 spots")
	case 5:
		fmt.Println("You can move 5 spots")

	case 6:
		fmt.Println("You can move 6 spots and roll the dice again")
	default:
		fmt.Println("what was that?")
	}
}
