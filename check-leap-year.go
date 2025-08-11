package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Leap Year Checker")
	fmt.Println("Enter a year to check if it's a leap year:")
	year, err := readInt(reader, "")
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
	if isLeapYear(year) {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}
