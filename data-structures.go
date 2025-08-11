package main

import "fmt"

func main() {
	fmt.Println("=== GO DATA STRUCTURES DEMO ===\n")

	fmt.Println("1. ARRAYS:")
	fmt.Println("-----------")

	var numbers [5]int
	fmt.Println("Empty array:", numbers)

	scores := [3]int{85, 92, 78}
	fmt.Println("Scores array:", scores)

	colors := [...]string{"red", "green", "blue", "yellow"}
	fmt.Println("Colors array:", colors)
	fmt.Println("Array length:", len(colors))

	fmt.Println("First color:", colors[0])
	fmt.Println("Last color:", colors[len(colors)-1])

	colors[1] = "purple"
	fmt.Println("Modified colors:", colors)

	fmt.Println("Iterating over colors:")
	for i, color := range colors {
		fmt.Printf("Index %d: %s\n", i, color)
	}

	fmt.Println()

	fmt.Println("2. SLICES (Dynamic Arrays):")
	fmt.Println("-----------------------------")

	var emptySlice []int
	fmt.Println("Empty slice:", emptySlice, "Length:", len(emptySlice), "Capacity:", cap(emptySlice))

	fruits := []string{"apple", "banana", "orange"}
	fmt.Println("Fruits slice:", fruits, "Length:", len(fruits), "Capacity:", cap(fruits))

	numbersArray := [5]int{1, 2, 3, 4, 5}
	sliceFromArray := numbersArray[1:4]
	fmt.Println("Slice from array:", sliceFromArray)

	dynamicSlice := make([]int, 3, 5)
	fmt.Println("Dynamic slice:", dynamicSlice, "Length:", len(dynamicSlice), "Capacity:", cap(dynamicSlice))

	fruits = append(fruits, "grape")
	fmt.Println("After append:", fruits)

	fruits = append(fruits, "mango", "kiwi")
	fmt.Println("After multiple append:", fruits)

	moreFruits := []string{"pineapple", "strawberry"}
	fruits = append(fruits, moreFruits...)
	fmt.Println("After appending slice:", fruits)

	fruits = fruits[:len(fruits)-1] // Remove last element
	fmt.Println("After removing last:", fruits)

	original := []int{1, 2, 3, 4, 5}
	copied := make([]int, len(original))
	copy(copied, original)
	fmt.Println("Original:", original)
	fmt.Println("Copied:", copied)

	copied[0] = 99
	fmt.Println("After modifying copied:", copied)
	fmt.Println("Original unchanged:", original)

	fmt.Println()

	fmt.Println("3. MAPS (Key-Value Pairs):")
	fmt.Println("----------------------------")

	var emptyMap map[string]int
	fmt.Println("Empty map:", emptyMap)

	person := make(map[string]string)
	person["name"] = "Alice"
	person["city"] = "New York"
	person["job"] = "Developer"
	fmt.Println("Person map:", person)

	scoresMap := map[string]int{
		"Alice":   85,
		"Bob":     92,
		"Charlie": 78,
		"David":   95,
	}
	fmt.Println("Scores map:", scoresMap)

	fmt.Println("Alice's score:", scoresMap["Alice"])

	score, exists := scoresMap["Eve"]
	if exists {
		fmt.Println("Eve's score:", score)
	} else {
		fmt.Println("Eve not found in scores")
	}

	scoresMap["Eve"] = 88
	fmt.Println("After adding Eve:", scoresMap)

	scoresMap["Alice"] = 90
	fmt.Println("After updating Alice:", scoresMap)

	delete(scoresMap, "Charlie")
	fmt.Println("After deleting Charlie:", scoresMap)

	fmt.Println("Iterating over scores:")
	for name, score := range scoresMap {
		fmt.Printf("%s: %d\n", name, score)
	}

	studentInfo := map[string]interface{}{
		"name":   "John",
		"age":    20,
		"grades": []int{85, 92, 78},
		"active": true,
		"height": 175.5,
	}
	fmt.Println("Student info:", studentInfo)

	fmt.Println()

	fmt.Println("4. PRACTICAL EXAMPLES:")
	fmt.Println("-----------------------")

	fmt.Println("Student Grade Tracker:")
	students := []string{"Alice", "Bob", "Charlie", "David"}
	grades := make(map[string][]int)

	grades["Alice"] = []int{85, 92, 78, 95}
	grades["Bob"] = []int{90, 88, 92, 87}
	grades["Charlie"] = []int{75, 82, 79, 85}
	grades["David"] = []int{95, 98, 92, 96}

	for _, student := range students {
		if studentGrades, exists := grades[student]; exists {
			sum := 0
			for _, grade := range studentGrades {
				sum += grade
			}
			average := float64(sum) / float64(len(studentGrades))
			fmt.Printf("%s: Average = %.2f (Grades: %v)\n", student, average, studentGrades)
		}
	}

	fmt.Println()

	fmt.Println("Inventory Management:")
	inventory := map[string]map[string]interface{}{
		"laptop": {
			"quantity": 10,
			"price":    999.99,
			"category": "electronics",
		},
		"book": {
			"quantity": 50,
			"price":    19.99,
			"category": "books",
		},
		"pen": {
			"quantity": 100,
			"price":    2.50,
			"category": "stationery",
		},
	}

	for item, details := range inventory {
		quantity := details["quantity"].(int)
		price := details["price"].(float64)
		category := details["category"].(string)
		totalValue := float64(quantity) * price
		fmt.Printf("%s (%s): %d units, $%.2f each, Total: $%.2f\n",
			item, category, quantity, price, totalValue)
	}

	fmt.Println()

	fmt.Println("Data Processing Pipeline:")

	rawData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var evenNumbers []int
	for _, num := range rawData {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}
	fmt.Println("Even numbers:", evenNumbers)

	var squaredNumbers []int
	for _, num := range evenNumbers {
		squaredNumbers = append(squaredNumbers, num*num)
	}
	fmt.Println("Squared numbers:", squaredNumbers)

	sum := 0
	for _, num := range squaredNumbers {
		sum += num
	}
	fmt.Println("Sum of squared even numbers:", sum)

	fmt.Println("\n=== END OF DATA STRUCTURES DEMO ===")
}
