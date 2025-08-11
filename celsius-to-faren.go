package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func celsiusToFarenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func farenheitToCelsius(farenheit float64) float64 {
	return (farenheit - 32) * 5 / 9
}

func readFloat(reader *bufio.Reader, prompt string) (float64, error) {
	for {
		raw, err := readLine(reader, prompt)
		if err != nil {
			return 0, err
		}
		value, err := strconv.ParseFloat(strings.TrimSpace(raw), 64)
		if err != nil {
			return 0, err
		}
		return value, nil
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
	fmt.Println("Temperature Converter")
	fmt.Println("Enter a temperature in Celsius:")
	celsius, err := readFloat(reader, "")
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
	fmt.Printf("%.2f°C is %.2f°F\n", celsius, celsiusToFarenheit(celsius))
	fmt.Println("Enter a temperature in Fahrenheit:")
	fahrenheit, err := readFloat(reader, "")
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
	fmt.Printf("%.2f°F is %.2f°C\n", fahrenheit, farenheitToCelsius(fahrenheit))
}
