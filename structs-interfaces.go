package main

import (
	"fmt"
	"math"
)


type Person struct {
	Name string
	Age  int
	City string
}


type Student struct {
	Person        
	StudentID     string
	Major         string
	GPA           float64
	IsGraduated   bool
}


	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}


func (p Person) Introduce() string {
	return fmt.Sprintf("Hi, I'm %s, %d years old from %s", p.Name, p.Age, p.City)
}

func (p *Person) HaveBirthday() {
	p.Age++
}


func (s Student) GetFullInfo() string {
	return fmt.Sprintf("Student: %s (ID: %s), Major: %s, GPA: %.2f", 
		s.Name, s.StudentID, s.Major, s.GPA)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}


func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}


type Shape interface {
	Area() float64
	Perimeter() float64
}

type Describable interface {
	GetDescription() string
}

type FinancialAccount interface {
	Describable
	GetBalance() float64
}

func (p Person) GetDescription() string {
	return fmt.Sprintf("Person named %s", p.Name)
}

func (s Student) GetDescription() string {
	return fmt.Sprintf("Student %s studying %s", s.Name, s.Major)
}


func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}


func Describe(d Describable) {
	fmt.Println(d.GetDescription())
}


func PrintAccountInfo(account FinancialAccount) {
	Describe(account)
	fmt.Printf("Current balance: $%.2f\n", account.GetBalance())
}


func PrintAnything(value interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

func ProcessValue(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case Person:
		fmt.Printf("Person: %s\n", v.Introduce())
	case Student:
		fmt.Printf("Student: %s\n", v.GetFullInfo())
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}


type BankAccount struct {
	AccountNumber string
	Owner         Person
	Balance       float64
	Transactions  []Transaction
}

type Transaction struct {
	Type      string  
	Amount    float64
	Timestamp string
}

func (ba *BankAccount) Deposit(amount float64) {
	ba.Balance += amount
	ba.Transactions = append(ba.Transactions, Transaction{
		Type:      "deposit",
		Amount:    amount,
		Timestamp: "2024-01-01 10:00:00",
	})
}

func (ba *BankAccount) Withdraw(amount float64) bool {
	if ba.Balance >= amount {
		ba.Balance -= amount
		ba.Transactions = append(ba.Transactions, Transaction{
			Type:      "withdrawal",
			Amount:    amount,
			Timestamp: "2024-01-01 10:00:00",
		})
		return true
	}
	return false
}

func (ba BankAccount) GetBalance() float64 {
	return ba.Balance
}

func (ba BankAccount) GetTransactionHistory() []Transaction {
	return ba.Transactions
}


func (ba BankAccount) GetDescription() string {
	return fmt.Sprintf("Bank account %s owned by %s with balance $%.2f", 
		ba.AccountNumber, ba.Owner.Name, ba.Balance)
}



func main() {
	fmt.Println("=== GO STRUCTS & INTERFACES DEMO ===\n")

	fmt.Println("1. BASIC STRUCTS:")
	fmt.Println("------------------")

	person1 := Person{
		Name: "Alice",
		Age:  25,
		City: "New York",
	}
	fmt.Println("Person:", person1)
	fmt.Println("Introduction:", person1.Introduce())

	student1 := Student{
		Person: Person{
			Name: "Bob",
			Age:  20,
			City: "Boston",
		},
		StudentID:   "S12345",
		Major:       "Computer Science",
		GPA:         3.8,
		IsGraduated: false,
	}
	fmt.Println("Student:", student1)
	fmt.Println("Student info:", student1.GetFullInfo())

	fmt.Printf("Student's name: %s, age: %d\n", student1.Name, student1.Age)

	fmt.Println()

	fmt.Println("2. METHODS & POINTER RECEIVERS:")
	fmt.Println("--------------------------------")

	fmt.Println("Before birthday:", person1.Age)
	person1.HaveBirthday()
	fmt.Println("After birthday:", person1.Age)

	fmt.Println()

	fmt.Println("3. INTERFACES & POLYMORPHISM:")
	fmt.Println("------------------------------")

	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	fmt.Println("Rectangle:")
	PrintShapeInfo(rect)

	fmt.Println("Circle:")
	PrintShapeInfo(circle)

	fmt.Println("\nDescribing objects:")
	Describe(person1)
	Describe(student1)

	fmt.Println()

	fmt.Println("4. TYPE ASSERTIONS:")
	fmt.Println("-------------------")

	describables := []Describable{person1, student1}

	for i, d := range describables {
		fmt.Printf("Item %d: %s\n", i+1, d.GetDescription())
		
		if person, ok := d.(Person); ok {
			fmt.Printf("  -> This is a Person: %s\n", person.Introduce())
		} else if student, ok := d.(Student); ok {
			fmt.Printf("  -> This is a Student: %s\n", student.GetFullInfo())
		}
	}

	fmt.Println()

	// 5. Practical example: Bank Account System
	fmt.Println("5. PRACTICAL EXAMPLE: BANK ACCOUNT SYSTEM:")
	fmt.Println("-------------------------------------------")

	// Create bank account
	account := BankAccount{
		AccountNumber: "ACC123456",
		Owner: Person{
			Name: "Charlie",
			Age:  30,
			City: "San Francisco",
		},
		Balance:      1000.0,
		Transactions: []Transaction{},
	}

	fmt.Println("Initial account info:")
	Describe(account)

	fmt.Println("\nPerforming transactions:")
	account.Deposit(500.0)
	fmt.Printf("After deposit: Balance = $%.2f\n", account.GetBalance())

	success := account.Withdraw(200.0)
	if success {
		fmt.Printf("After withdrawal: Balance = $%.2f\n", account.GetBalance())
	} else {
		fmt.Println("Withdrawal failed - insufficient funds")
	}

	success = account.Withdraw(2000.0)
	if !success {
		fmt.Println("Withdrawal failed - insufficient funds")
	}

	fmt.Println("\nTransaction history:")
	for i, transaction := range account.GetTransactionHistory() {
		fmt.Printf("  %d. %s: $%.2f at %s\n", 
			i+1, transaction.Type, transaction.Amount, transaction.Timestamp)
	}

	fmt.Println()

	fmt.Println("6. INTERFACE COMPOSITION:")
	fmt.Println("--------------------------")

	PrintAccountInfo(account)

	fmt.Println()

	fmt.Println("7. EMPTY INTERFACE:")
	fmt.Println("-------------------")

	PrintAnything(42)
	PrintAnything("Hello, Go!")
	PrintAnything(person1)
	PrintAnything([]int{1, 2, 3})

	fmt.Println()

	fmt.Println("8. TYPE SWITCHES:")
	fmt.Println("-----------------")

	ProcessValue(100)
	ProcessValue("Go Programming")
	ProcessValue(person1)
	ProcessValue(student1)
	ProcessValue(3.14)

	fmt.Println("\n=== END OF STRUCTS & INTERFACES DEMO ===")
} 