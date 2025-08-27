package main

import "fmt"

func main() {
	fmt.Println("Exercise 4: Maps")
	fmt.Println("================")

	// 1. Create and initialize maps in different ways
	// Using make()
	studentGrades := make(map[string]string)
	studentGrades["Alice"] = "A"
	studentGrades["Bob"] = "B"
	studentGrades["Charlie"] = "A"
	studentGrades["Diana"] = "C"
	studentGrades["Eve"] = "B"

	// Using literal syntax
	colors := map[string]string{
		"apple":  "red",
		"banana": "yellow",
		"cherry": "red",
	}

	// Using map[T]T{} syntax
	config := map[string]interface{}{
		"debug":   true,
		"port":    8080,
		"timeout": 30,
	}

	// 2. Print student grades
	fmt.Println("Student Grades:")
	for name, grade := range studentGrades {
		fmt.Printf("%s: %s\n", name, grade)
	}

	// 3. Grade distribution
	gradeDistribution := make(map[string]int)
	for _, grade := range studentGrades {
		gradeDistribution[grade]++
	}

	fmt.Println("\nGrade Distribution:")
	for grade, count := range gradeDistribution {
		fmt.Printf("%s: %d students\n", grade, count)
	}

	// 4. Nested maps for user database
	userDatabase := map[int]map[string]interface{}{
		1: {
			"Name":  "John Doe",
			"Age":   25,
			"Email": "john@example.com",
		},
		2: {
			"Name":  "Jane Smith",
			"Age":   30,
			"Email": "jane@example.com",
		},
	}

	fmt.Println("\nUser Database:")
	for id, user := range userDatabase {
		fmt.Printf("User %d:\n", id)
		fmt.Printf("  Name: %s\n", user["Name"])
		fmt.Printf("  Age: %d\n", user["Age"])
		fmt.Printf("  Email: %s\n", user["Email"])
	}

	// 5. Word frequency counter
	text := "hello world hello go world"
	words := []string{"hello", "world", "hello", "go", "world"}
	wordFrequency := make(map[string]int)

	for _, word := range words {
		wordFrequency[word]++
	}

	fmt.Println("\nWord Frequency Analysis:")
	fmt.Printf("Text: \"%s\"\n", text)
	for word, count := range wordFrequency {
		fmt.Printf("%s: %d\n", word, count)
	}

	// 6. Configuration settings
	appConfig := map[string]map[string]interface{}{
		"Server": {
			"Host": "localhost",
			"Port": 8080,
		},
		"Database": {
			"Host": "db.example.com",
			"Port": 5432,
			"Name": "myapp",
		},
	}

	fmt.Println("\nConfiguration Settings:")
	for section, settings := range appConfig {
		fmt.Printf("%s:\n", section)
		for key, value := range settings {
			fmt.Printf("  %s: %v\n", key, value)
		}
	}

	// 7. Map operations demonstration
	fmt.Println("\nMap Operations:")
	fmt.Printf("Original map: %v\n", colors)

	// Adding a new key-value pair
	colors["orange"] = "orange"
	fmt.Printf("After adding orange: %v\n", colors)

	// Updating an existing value
	colors["apple"] = "green"
	fmt.Printf("After updating apple: %v\n", colors)

	// Deleting a key
	delete(colors, "banana")
	fmt.Printf("After deleting banana: %v\n", colors)

	// 8. Key existence check
	fmt.Println("\nKey existence check:")
	if value, exists := colors["apple"]; exists {
		fmt.Printf("Key 'apple' exists: true (value: %s)\n", value)
	} else {
		fmt.Println("Key 'apple' exists: false")
	}

	if _, exists := colors["grape"]; exists {
		fmt.Println("Key 'grape' exists: true")
	} else {
		fmt.Println("Key 'grape' exists: false")
	}

	// 9. Map with different value types
	mixedMap := map[string]interface{}{
		"string": "hello",
		"int":    42,
		"float":  3.14,
		"bool":   true,
		"slice":  []int{1, 2, 3},
		"nested": map[string]string{"key": "value"},
	}

	fmt.Println("\nMixed type map:")
	for key, value := range mixedMap {
		fmt.Printf("  %s: %v (%T)\n", key, value, value)
	}

	// 10. Map iteration order (not guaranteed)
	fmt.Println("\nMap iteration (order not guaranteed):")
	for key, value := range studentGrades {
		fmt.Printf("  %s -> %s\n", key, value)
	}
}
