package main

import (
	"fmt"
	"testing"
)

// TestAdd tests the Add function
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -1, -2, -3},
		{"mixed numbers", -1, 5, 4},
		{"zero", 0, 10, 10},
		{"large numbers", 1000, 2000, 3000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestMultiply tests the Multiply function
func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 4, 5, 20},
		{"negative numbers", -2, 3, -6},
		{"zero", 0, 10, 0},
		{"one", 1, 100, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestDivide tests the Divide function
func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a, b        int
		expected    float64
		expectError bool
	}{
		{"normal division", 10, 2, 5.0, false},
		{"division by zero", 10, 0, 0, true},
		{"negative division", -10, 2, -5.0, false},
		{"fraction result", 5, 2, 2.5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)

			if tt.expectError {
				if err == nil {
					t.Errorf("Divide(%d, %d) expected error but got none", tt.a, tt.b)
				}
			} else {
				if err != nil {
					t.Errorf("Divide(%d, %d) unexpected error: %v", tt.a, tt.b, err)
				}
				if result != tt.expected {
					t.Errorf("Divide(%d, %d) = %f; want %f", tt.a, tt.b, result, tt.expected)
				}
			}
		})
	}
}

// TestStringReverse tests the StringReverse function
func TestStringReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple string", "hello", "olleh"},
		{"palindrome", "racecar", "racecar"},
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"unicode", "café", "éfac"},
		{"numbers", "12345", "54321"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringReverse(tt.input)
			if result != tt.expected {
				t.Errorf("StringReverse(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestStringReverseTable tests StringReverse with table-driven tests
func TestStringReverseTable(t *testing.T) {
	testCases := map[string]string{
		"hello":       "olleh",
		"world":       "dlrow",
		"golang":      "gnalog",
		"":            "",
		"a":           "a",
		"123":         "321",
		"hello world": "dlrow olleh",
	}

	for input, expected := range testCases {
		t.Run(input, func(t *testing.T) {
			result := StringReverse(input)
			if result != expected {
				t.Errorf("StringReverse(%q) = %q; want %q", input, result, expected)
			}
		})
	}
}

// TestCalculator tests the Calculator struct methods
func TestCalculator(t *testing.T) {
	calc := &Calculator{}

	t.Run("Add", func(t *testing.T) {
		result := calc.Add(5, 3)
		expected := 8
		if result != expected {
			t.Errorf("Calculator.Add(5, 3) = %d; want %d", result, expected)
		}
	})

	t.Run("Subtract", func(t *testing.T) {
		result := calc.Subtract(10, 4)
		expected := 6
		if result != expected {
			t.Errorf("Calculator.Subtract(10, 4) = %d; want %d", result, expected)
		}
	})

	t.Run("Multiply", func(t *testing.T) {
		result := calc.Multiply(6, 7)
		expected := 42
		if result != expected {
			t.Errorf("Calculator.Multiply(6, 7) = %d; want %d", result, expected)
		}
	})

	t.Run("Divide", func(t *testing.T) {
		result, err := calc.Divide(20, 5)
		if err != nil {
			t.Errorf("Calculator.Divide(20, 5) unexpected error: %v", err)
		}
		expected := 4.0
		if result != expected {
			t.Errorf("Calculator.Divide(20, 5) = %f; want %f", result, expected)
		}
	})

	t.Run("Divide by zero", func(t *testing.T) {
		_, err := calc.Divide(10, 0)
		if err == nil {
			t.Error("Calculator.Divide(10, 0) expected error but got none")
		}
	})

	t.Run("History", func(t *testing.T) {
		history := calc.GetHistory()
		if len(history) == 0 {
			t.Error("Calculator history should not be empty after operations")
		}
	})

	t.Run("Clear history", func(t *testing.T) {
		calc.ClearHistory()
		history := calc.GetHistory()
		if len(history) != 0 {
			t.Error("Calculator history should be empty after clearing")
		}
	})
}

// TestIsPrime tests the IsPrime function
func TestIsPrime(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"prime number", 7, true},
		{"not prime", 4, false},
		{"zero", 0, false},
		{"one", 1, false},
		{"two", 2, true},
		{"large prime", 97, true},
		{"large not prime", 100, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPrime(tt.input)
			if result != tt.expected {
				t.Errorf("IsPrime(%d) = %t; want %t", tt.input, result, tt.expected)
			}
		})
	}
}

// TestFibonacci tests the Fibonacci function
func TestFibonacci(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"zero", 0, 0},
		{"one", 1, 1},
		{"two", 2, 1},
		{"three", 3, 2},
		{"four", 4, 3},
		{"five", 5, 5},
		{"six", 6, 8},
		{"seven", 7, 13},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Fibonacci(tt.input)
			if result != tt.expected {
				t.Errorf("Fibonacci(%d) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(10, 20)
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(10, 20)
	}
}

func BenchmarkStringReverse(b *testing.B) {
	input := "hello world this is a test string"
	for i := 0; i < b.N; i++ {
		StringReverse(input)
	}
}

func BenchmarkCalculator(b *testing.B) {
	calc := &Calculator{}
	for i := 0; i < b.N; i++ {
		calc.Add(i, i+1)
	}
}

// Example tests
func ExampleAdd() {
	result := Add(2, 3)
	fmt.Println(result)
	// Output: 5
}

func ExampleStringReverse() {
	result := StringReverse("hello")
	fmt.Println(result)
	// Output: olleh
}

func ExampleCalculator_Add() {
	calc := &Calculator{}
	result := calc.Add(5, 3)
	fmt.Println(result)
	// Output: 8
}
