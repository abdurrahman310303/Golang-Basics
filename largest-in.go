package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func largestInArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	largest := arr[0]

	for _, v := range arr {
		if v > largest {
			largest = v
		}
	}
	return largest
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
	fmt.Println("Enter a list of integers separated by spaces:")
	raw, err := readLine(reader, "")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	numbers := strings.Fields(raw)
	arr := make([]int, len(numbers))
	valid := true

	for i, num := range numbers {
		v, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Invalid integer:", num)
			valid = false
			break
		}
		arr[i] = v
	}

	if valid && len(arr) > 0 {
		largest := largestInArray(arr)
		fmt.Printf("The largest number is: %d\n", largest)
	} else {
		fmt.Println("Please enter a valid list of integers")
	}
}
