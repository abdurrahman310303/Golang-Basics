package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// ---- Math functions ----

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func pow(a, b float64) float64 {
	return math.Pow(a, b)
}

func sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("cannot take square root of a negative number")
	}
	return math.Sqrt(a), nil
}

func factorial(n int) (uint64, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial undefined for negative numbers")
	}
	// Limit to 20 to avoid uint64 overflow
	if n > 20 {
		return 0, fmt.Errorf("n too large (max 20 for uint64), got %d", n)
	}
	var result uint64 = 1
	for i := 2; i <= n; i++ {
		result *= uint64(i)
	}
	return result, nil
}

func average(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("cannot compute average of empty list")
	}
	sum := 0.0
	for _, v := range nums {
		sum += v
	}
	return sum / float64(len(nums)), nil
}

// ---- Input helpers ----

func readLine(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

func readFloat(reader *bufio.Reader, prompt string) (float64, error) {
	for {
		raw, err := readLine(reader, prompt)
		if err != nil {
			return 0, err
		}
		v, err := strconv.ParseFloat(strings.TrimSpace(raw), 64)
		if err != nil {
			fmt.Println("Invalid number, try again.")
			continue
		}
		return v, nil
	}
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

func readCSVFloats(reader *bufio.Reader, prompt string) ([]float64, error) {
	for {
		raw, err := readLine(reader, prompt)
		if err != nil {
			return nil, err
		}
		raw = strings.TrimSpace(raw)
		if raw == "" {
			fmt.Println("Please enter at least one number.")
			continue
		}
		parts := strings.Split(raw, ",")
		var nums []float64
		ok := true
		for _, p := range parts {
			p = strings.TrimSpace(p)
			v, err := strconv.ParseFloat(p, 64)
			if err != nil {
				fmt.Println("Invalid list, please enter comma-separated numbers. Example: 1, 2.5, 3")
				ok = false
				break
			}
			nums = append(nums, v)
		}
		if ok {
			return nums, nil
		}
	}
}

// ---- Menu ----

func printMenu() {
	fmt.Println("\n=== CLI Calculator ===")
	fmt.Println("1) Add (a + b)")
	fmt.Println("2) Subtract (a - b)")
	fmt.Println("3) Multiply (a * b)")
	fmt.Println("4) Divide (a / b)")
	fmt.Println("5) Power (a ^ b)")
	fmt.Println("6) Square Root (√a)")
	fmt.Println("7) Factorial (n!)")
	fmt.Println("8) Average of list")
	fmt.Println("9) Show history")
	fmt.Println("0) Exit")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var history []string

	for {
		printMenu()
		choice, err := readLine(reader, "Choose an option: ")
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		switch choice {
		case "1":
			a, _ := readFloat(reader, "Enter a: ")
			b, _ := readFloat(reader, "Enter b: ")
			res := add(a, b)
			line := fmt.Sprintf("%.6g + %.6g = %.6g", a, b, res)
			fmt.Println(line)
			history = append(history, line)

		case "2":
			a, _ := readFloat(reader, "Enter a: ")
			b, _ := readFloat(reader, "Enter b: ")
			res := sub(a, b)
			line := fmt.Sprintf("%.6g - %.6g = %.6g", a, b, res)
			fmt.Println(line)
			history = append(history, line)

		case "3":
			a, _ := readFloat(reader, "Enter a: ")
			b, _ := readFloat(reader, "Enter b: ")
			res := mul(a, b)
			line := fmt.Sprintf("%.6g * %.6g = %.6g", a, b, res)
			fmt.Println(line)
			history = append(history, line)

		case "4":
			a, _ := readFloat(reader, "Enter a: ")
			b, _ := readFloat(reader, "Enter b: ")
			res, err := div(a, b)
			if err != nil {
				fmt.Println("Error:", err)
				history = append(history, fmt.Sprintf("DIV ERROR: %v", err))
				break
			}
			line := fmt.Sprintf("%.6g / %.6g = %.6g", a, b, res)
			fmt.Println(line)
			history = append(history, line)

		case "5":
			a, _ := readFloat(reader, "Enter base a: ")
			b, _ := readFloat(reader, "Enter exponent b: ")
			res := pow(a, b)
			line := fmt.Sprintf("%.6g ^ %.6g = %.6g", a, b, res)
			fmt.Println(line)
			history = append(history, line)

		case "6":
			a, _ := readFloat(reader, "Enter a: ")
			res, err := sqrt(a)
			if err != nil {
				fmt.Println("Error:", err)
				history = append(history, fmt.Sprintf("SQRT ERROR: %v", err))
				break
			}
			line := fmt.Sprintf("√(%.6g) = %.6g", a, res)
			fmt.Println(line)
			history = append(history, line)

		case "7":
			n, _ := readInt(reader, "Enter n (0..20): ")
			res, err := factorial(n)
			if err != nil {
				fmt.Println("Error:", err)
				history = append(history, fmt.Sprintf("FACT ERROR: %v", err))
				break
			}
			line := fmt.Sprintf("%d! = %d", n, res)
			fmt.Println(line)
			history = append(history, line)

		case "8":
			nums, _ := readCSVFloats(reader, "Enter numbers (comma-separated): ")
			res, err := average(nums)
			if err != nil {
				fmt.Println("Error:", err)
				history = append(history, fmt.Sprintf("AVG ERROR: %v", err))
				break
			}
			line := fmt.Sprintf("avg(%v) = %.6g", nums, res)
			fmt.Println(line)
			history = append(history, line)

		case "9":
			if len(history) == 0 {
				fmt.Println("History is empty.")
				break
			}
			fmt.Println("=== History ===")
			for i, h := range history {
				fmt.Printf("%2d) %s\n", i+1, h)
			}

		case "0":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}
