package main

import (
	"fmt"
	"math"
)

// Calculator struct for testing
type Calculator struct {
	History []string
}

// Add function
func Add(a, b int) int {
	return a + b
}

// Multiply function
func Multiply(a, b int) int {
	return a * b
}

// Divide function with error handling
func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return float64(a) / float64(b), nil
}

// String reverse function
func StringReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Calculator methods
func (c *Calculator) Add(a, b int) int {
	result := a + b
	c.History = append(c.History, fmt.Sprintf("%d + %d = %d", a, b, result))
	return result
}

func (c *Calculator) Subtract(a, b int) int {
	result := a - b
	c.History = append(c.History, fmt.Sprintf("%d - %d = %d", a, b, result))
	return result
}

func (c *Calculator) Multiply(a, b int) int {
	result := a * b
	c.History = append(c.History, fmt.Sprintf("%d * %d = %d", a, b, result))
	return result
}

func (c *Calculator) Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	result := float64(a) / float64(b)
	c.History = append(c.History, fmt.Sprintf("%d / %d = %.2f", a, b, result))
	return result, nil
}

func (c *Calculator) GetHistory() []string {
	return c.History
}

func (c *Calculator) ClearHistory() {
	c.History = []string{}
}

// Math utility functions
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	fmt.Println("Exercise 9: Unit Testing")
	fmt.Println("=========================")

	// Demonstrate the functions that will be tested
	fmt.Println("Testing Add function:")
	fmt.Printf("Add(2, 3) = %d\n", Add(2, 3))
	fmt.Printf("Add(-1, 5) = %d\n", Add(-1, 5))

	fmt.Println("\nTesting Multiply function:")
	fmt.Printf("Multiply(4, 5) = %d\n", Multiply(4, 5))
	fmt.Printf("Multiply(0, 10) = %d\n", Multiply(0, 10))

	fmt.Println("\nTesting Divide function:")
	if result, err := Divide(10, 2); err == nil {
		fmt.Printf("Divide(10, 2) = %.2f\n", result)
	}
	if result, err := Divide(10, 0); err != nil {
		fmt.Printf("Divide(10, 0) = Error: %v\n", err)
	}

	fmt.Println("\nTesting StringReverse function:")
	fmt.Printf("StringReverse(\"hello\") = %s\n", StringReverse("hello"))
	fmt.Printf("StringReverse(\"racecar\") = %s\n", StringReverse("racecar"))

	fmt.Println("\nTesting Calculator struct:")
	calc := &Calculator{}
	fmt.Printf("Calculator.Add(5, 3) = %d\n", calc.Add(5, 3))
	fmt.Printf("Calculator.Subtract(10, 4) = %d\n", calc.Subtract(10, 4))
	fmt.Printf("Calculator.Multiply(6, 7) = %d\n", calc.Multiply(6, 7))
	if result, err := calc.Divide(20, 5); err == nil {
		fmt.Printf("Calculator.Divide(20, 5) = %.2f\n", result)
	} else {
		fmt.Printf("Calculator.Divide(20, 5) = Error: %v\n", err)
	}

	fmt.Println("\nCalculator History:")
	for _, entry := range calc.GetHistory() {
		fmt.Printf("  %s\n", entry)
	}

	fmt.Println("\nTesting utility functions:")
	fmt.Printf("IsPrime(7) = %t\n", IsPrime(7))
	fmt.Printf("IsPrime(4) = %t\n", IsPrime(4))
	fmt.Printf("Fibonacci(8) = %d\n", Fibonacci(8))

	fmt.Println("\nRun 'go test' to execute the test suite!")
}
