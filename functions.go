package main

import "fmt"

// Function with no parameters and no return value
func sayHello() {
    fmt.Println("Hello from a function!")
}

// Function with parameters but no return value
func greet(name string) {
    fmt.Println("Hello,", name)
}

// Function with parameters and return value
func add(a int, b int) int {
    return a + b
}

// Function with multiple return values
func divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    // Call function with no parameters
    sayHello()
    
    // Call function with parameters
    greet("Alice")
    greet("Bob")
    
    // Call function with return value
    result := add(5, 3)
    fmt.Println("5 + 3 =", result)
    
    // Call function with multiple return values
    quotient, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("10 / 2 =", quotient)
    }
    
    // Try division by zero
    _, err2 := divide(10, 0)
    if err2 != nil {
        fmt.Println("Error:", err2)
    }
} 