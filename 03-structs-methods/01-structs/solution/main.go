package main

import (
	"encoding/json"
	"fmt"
	"math"
)

// Person struct with basic fields
type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
}

// Methods for Person struct
func (p Person) GetFullInfo() string {
	return fmt.Sprintf("%s (%d) - %s", p.Name, p.Age, p.Email)
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Methods for Rectangle struct
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) IsSquare() bool {
	return r.Width == r.Height
}

// BankAccount struct with transaction history
type BankAccount struct {
	AccountNumber string
	Balance       float64
	Transactions  []Transaction
}

type Transaction struct {
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

// Methods for BankAccount struct
func (ba *BankAccount) Deposit(amount float64) {
	ba.Balance += amount
	ba.Transactions = append(ba.Transactions, Transaction{
		Type:   "Deposit",
		Amount: amount,
	})
}

func (ba *BankAccount) Withdraw(amount float64) bool {
	if ba.Balance >= amount {
		ba.Balance -= amount
		ba.Transactions = append(ba.Transactions, Transaction{
			Type:   "Withdrawal",
			Amount: amount,
		})
		return true
	}
	return false
}

func (ba BankAccount) GetTransactionHistory() []Transaction {
	return ba.Transactions
}

// Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle struct implementing Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle already has Perimeter method defined above

// CustomPerson with custom JSON marshaling
type CustomPerson struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Salary float64 `json:"salary"`
}

func (cp CustomPerson) MarshalJSON() ([]byte, error) {
	type Alias CustomPerson
	return json.Marshal(&struct {
		Alias
		FormattedSalary string `json:"formatted_salary"`
	}{
		Alias:           Alias(cp),
		FormattedSalary: fmt.Sprintf("$%.2f", cp.Salary),
	})
}

func main() {
	fmt.Println("Exercise 5: Structs and Methods")
	fmt.Println("================================")

	// 1. Person struct with methods
	person := Person{
		Name:   "John Doe",
		Age:    30,
		Email:  "john@example.com",
		Active: true,
	}

	fmt.Println("Person Information:")
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("Email: %s\n", person.Email)
	fmt.Printf("Full Info: %s\n", person.GetFullInfo())

	// 2. Rectangle struct with methods
	rect := Rectangle{Width: 10, Height: 5}

	fmt.Println("\nRectangle Operations:")
	fmt.Printf("Rectangle: Width=%.0f, Height=%.0f\n", rect.Width, rect.Height)
	fmt.Printf("Area: %.0f\n", rect.Area())
	fmt.Printf("Perimeter: %.0f\n", rect.Perimeter())
	fmt.Printf("Is Square: %t\n", rect.IsSquare())

	// 3. BankAccount struct with transaction methods
	account := &BankAccount{
		AccountNumber: "12345",
		Balance:       1000.0,
		Transactions:  []Transaction{},
	}

	fmt.Println("\nBank Account Operations:")
	fmt.Printf("Account: %s\n", account.AccountNumber)
	fmt.Printf("Balance: $%.2f\n", account.Balance)

	account.Deposit(500.0)
	fmt.Printf("After deposit $500: $%.2f\n", account.Balance)

	account.Withdraw(200.0)
	fmt.Printf("After withdrawal $200: $%.2f\n", account.Balance)

	fmt.Println("Transaction History:")
	for _, tx := range account.GetTransactionHistory() {
		fmt.Printf("  %s: $%.2f\n", tx.Type, tx.Amount)
	}

	// 4. Shape interface demonstration
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 5}

	shapes := []Shape{circle, rectangle}

	fmt.Println("\nShape Interface:")
	for _, shape := range shapes {
		switch s := shape.(type) {
		case Circle:
			fmt.Printf("Circle (radius=%.0f): Area=%.2f, Perimeter=%.2f\n", s.Radius, s.Area(), s.Perimeter())
		case Rectangle:
			fmt.Printf("Rectangle (%.0fx%.0f): Area=%.2f, Perimeter=%.2f\n", s.Width, s.Height, s.Area(), s.Perimeter())
		}
	}

	// 5. JSON serialization with struct tags
	person2 := Person{
		Name:   "Jane Doe",
		Age:    25,
		Email:  "jane@example.com",
		Active: true,
	}

	jsonData, err := json.Marshal(person2)
	if err == nil {
		fmt.Println("\nJSON Serialization:")
		fmt.Println(string(jsonData))
	}

	// 6. Custom JSON marshaling
	customPerson := CustomPerson{
		Name:   "Bob Smith",
		Age:    35,
		Salary: 75000.50,
	}

	customJSON, err := json.Marshal(customPerson)
	if err == nil {
		fmt.Println("\nCustom JSON Marshaling:")
		fmt.Println(string(customJSON))
	}

	// 7. Struct composition
	type Employee struct {
		Person
		EmployeeID string
		Department string
	}

	employee := Employee{
		Person: Person{
			Name:   "Alice Johnson",
			Age:    28,
			Email:  "alice@company.com",
			Active: true,
		},
		EmployeeID: "EMP001",
		Department: "Engineering",
	}

	fmt.Println("\nStruct Composition:")
	fmt.Printf("Employee: %s\n", employee.Name)
	fmt.Printf("ID: %s\n", employee.EmployeeID)
	fmt.Printf("Department: %s\n", employee.Department)
	fmt.Printf("Is Adult: %t\n", employee.IsAdult()) // Inherited method

	// 8. Pointer receivers vs value receivers
	fmt.Println("\nPointer vs Value Receivers:")

	// Value receiver (doesn't modify original)
	rect2 := Rectangle{Width: 5, Height: 5}
	fmt.Printf("Original rectangle: %+v\n", rect2)

	// This would be a value receiver method call
	// (but we don't have one in this example)

	// Pointer receiver (can modify original)
	account2 := &BankAccount{Balance: 100}
	fmt.Printf("Original balance: $%.2f\n", account2.Balance)
	account2.Deposit(50)
	fmt.Printf("After deposit: $%.2f\n", account2.Balance)
}
