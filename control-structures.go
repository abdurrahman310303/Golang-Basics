package main

import "fmt"

func main() {


    age := 18
    if age >= 18 {
        fmt.Println("You are an adult!")
    }

    score := 85
    if score >= 90 {
        fmt.Println("Grade: A")
    } else if score >= 80 {
        fmt.Println("Grade: B")
    } else if score >= 70 {
        fmt.Println("Grade: C")
    } else {
        fmt.Println("Grade: F")
    }

    if temperature := 25; temperature > 30 {
        fmt.Println("It's hot outside!")
    } else if temperature > 20 {
        fmt.Println("It's nice weather!")
    } else {
        fmt.Println("It's cold outside!")
    }

    fmt.Println()


    fmt.Println("Traditional for loop:")
    for i := 0; i < 5; i++ {
        fmt.Printf("Count: %d\n", i)
    }

 
    fmt.Println("\nWhile-style loop:")
    j := 0
    for j < 3 {
        fmt.Printf("While count: %d\n", j)
        j++
    }


    fmt.Println("\nInfinite loop with break:")
    k := 0
    for {
        if k >= 3 {
            break // Exit the loop
        }
        fmt.Printf("Infinite loop count: %d\n", k)
        k++
    }

    // Loop with continue
    fmt.Println("\nLoop with continue (skip even numbers):")
    for i := 0; i < 10; i++ {
        if i%2 == 0 {
            continue // Skip to next iteration
        }
        fmt.Printf("Odd number: %d\n", i)
    }

    fmt.Println("\nRange loop with slice:")
    fruits := []string{"apple", "banana", "orange", "grape"}
    for index, fruit := range fruits {
        fmt.Printf("Index %d: %s\n", index, fruit)
    }


    fmt.Println("\nRange loop with map:")
    person := map[string]string{
        "name": "Alice",
        "city": "New York",
        "job":  "Developer",
    }
    for key, value := range person {
        fmt.Printf("%s: %s\n", key, value)
    }

    // Range loop ignoring index/value
    fmt.Println("\nRange loop ignoring index:")
    for _, fruit := range fruits {
        fmt.Printf("Fruit: %s\n", fruit)
    }

    fmt.Println()


    fmt.Println("3. SWITCH STATEMENTS:")
    fmt.Println("----------------------")

    day := "Monday"
    switch day {
    case "Monday":
        fmt.Println("Start of work week")
    case "Tuesday", "Wednesday", "Thursday":
        fmt.Println("Mid week")
    case "Friday":
        fmt.Println("TGIF!")
    case "Saturday", "Sunday":
        fmt.Println("Weekend!")
    default:
        fmt.Println("Invalid day")
    }

    fmt.Println("\nSwitch with expression:")
    grade := 85
    switch {
    case grade >= 90:
        fmt.Println("Excellent!")
    case grade >= 80:
        fmt.Println("Good job!")
    case grade >= 70:
        fmt.Println("Passing")
    default:
        fmt.Println("Needs improvement")
    }

    // Switch with initialization
    fmt.Println("\nSwitch with initialization:")
    switch hour := 14; {
    case hour < 12:
        fmt.Println("Good morning!")
    case hour < 17:
        fmt.Println("Good afternoon!")
    default:
        fmt.Println("Good evening!")
    }

    // Switch with fallthrough
    fmt.Println("\nSwitch with fallthrough:")
    number := 2
    switch number {
    case 1:
        fmt.Println("One")
        fallthrough // Continue to next case
    case 2:
        fmt.Println("Two")
        fallthrough
    case 3:
        fmt.Println("Three")
    default:
        fmt.Println("Other number")
    }

    fmt.Println()

    // ========================================
    // 4. PRACTICAL EXAMPLE: NUMBER GUESSING GAME
    // ========================================
    fmt.Println("4. PRACTICAL EXAMPLE: NUMBER ANALYSIS")
    fmt.Println("--------------------------------------")

    numbers := []int{15, 7, 22, 3, 18, 9, 12, 6}
    
    for _, num := range numbers {
        fmt.Printf("Analyzing number %d: ", num)
        
        // Check if number is even or odd
        if num%2 == 0 {
            fmt.Print("Even, ")
        } else {
            fmt.Print("Odd, ")
        }
        
        // Check number range
        switch {
        case num < 10:
            fmt.Print("Single digit")
        case num < 20:
            fmt.Print("Teens")
        case num < 100:
            fmt.Print("Double digit")
        default:
            fmt.Print("Large number")
        }
        
        // Check if it's a special number
        switch num {
        case 7:
            fmt.Print(", Lucky number!")
        case 12:
            fmt.Print(", Dozen!")
        case 22:
            fmt.Print(", Double digits!")
        }
        
        fmt.Println()
    }

    fmt.Println("\n=== END OF CONTROL STRUCTURES DEMO ===")
} 