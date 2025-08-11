package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isEven(n int) bool {
	return n%2 == 0
}

func isOdd(n int) bool {
	return n%2 != 0
}

func readInt(reader *bufio.Reader, prompt string) (int, error) {

	for {
		raw, err := readLine(reader, prompt)
		if err != nil {
			return 0, err
		}
		v, err := strconv.Atoi(strings.TrimSpace(raw))
		if err != nil {
			fmt.Println("Invalid integer, try again.")
			continue
		}
		return v, nil
	}
}

func readLine(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

func main() {
	fmt.Println("Even or Odd Checker")
	fmt.Println("Enter a number to check if it's even or odd")

	reader := bufio.NewReader(os.Stdin)

	number, err := readInt(reader, "Enter a number: ")
	if err != nil {
		fmt.Println("Error reading input:", err)
	}

	if isEven(number) {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}
}
