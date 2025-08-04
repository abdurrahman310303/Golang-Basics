package main

import "fmt"

func sayHello() {
    fmt.Println("Hello from a function!")
}

func greet(name string) {
    fmt.Println("Hello,", name)
}

func add(a int, b int) int {
    return a + b
}

func divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    sayHello()
    
    greet("Alice")
    greet("Bob")
    
    result := add(5, 3)
    fmt.Println("5 + 3 =", result)
    
    quotient, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("10 / 2 =", quotient)
    }
    
    _, err2 := divide(10, 0)
    if err2 != nil {
        fmt.Println("Error:", err2)
    }
} 