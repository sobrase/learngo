package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Person struct with tags for demonstration
type Person struct {
	Name   string `json:"name" validate:"required"`
	Age    int    `json:"age" validate:"min:18"`
	Email  string `json:"email" validate:"email"`
	Active bool   `json:"active"`
}

// Config struct for configuration loading
type Config struct {
	Server struct {
		Host string `config:"host" default:"localhost"`
		Port int    `config:"port" default:"8080"`
	} `config:"server"`
	Database struct {
		Host string `config:"host" default:"db.example.com"`
		Port int    `config:"port" default:"5432"`
		Name string `config:"name" default:"myapp"`
	} `config:"database"`
}

// Methods for Person
func (p Person) GetFullName() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// Custom validator
func validateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// Struct inspector using reflection
func inspectStruct(v interface{}) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	fmt.Printf("Type: %s\n", typ)
	fmt.Printf("Kind: %s\n", val.Kind())
	fmt.Printf("Num fields: %d\n\n", val.NumField())

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		fmt.Printf("Field %d: %s (%s)\n", i, fieldType.Name, field.Type())
		fmt.Printf("  Tag: %s\n", fieldType.Tag)
		fmt.Printf("  Value: %v\n", field.Interface())
		fmt.Printf("  Can set: %t\n\n", field.CanSet())
	}
}

// JSON validator using reflection
func validateStruct(v interface{}) error {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	fmt.Printf("Validating %s struct...\n", typ.Name())

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// Get validation tag
		validateTag := fieldType.Tag.Get("validate")
		if validateTag == "" {
			continue
		}

		// Parse validation rules
		rules := strings.Split(validateTag, ",")
		for _, rule := range rules {
			rule = strings.TrimSpace(rule)

			switch {
			case rule == "required":
				if field.Kind() == reflect.String && field.String() == "" {
					return fmt.Errorf("field %s is required", fieldType.Name)
				}
				fmt.Printf("✓ %s: required field present\n", fieldType.Name)

			case strings.HasPrefix(rule, "min:"):
				minStr := strings.TrimPrefix(rule, "min:")
				min, _ := strconv.Atoi(minStr)
				if field.Kind() == reflect.Int {
					if field.Int() < int64(min) {
						return fmt.Errorf("field %s must be at least %d", fieldType.Name, min)
					}
					fmt.Printf("✓ %s: %d >= %d (min requirement met)\n", fieldType.Name, field.Int(), min)
				}

			case rule == "email":
				if field.Kind() == reflect.String {
					if !validateEmail(field.String()) {
						return fmt.Errorf("field %s must be a valid email", fieldType.Name)
					}
					fmt.Printf("✓ %s: valid email format\n", fieldType.Name)
				}
			}
		}
	}

	return nil
}

// Dynamic method calls using reflection
func callMethod(v interface{}, methodName string, args ...interface{}) ([]reflect.Value, error) {
	val := reflect.ValueOf(v)
	method := val.MethodByName(methodName)

	if !method.IsValid() {
		return nil, fmt.Errorf("method %s not found", methodName)
	}

	// Convert args to reflect.Value slice
	var reflectArgs []reflect.Value
	for _, arg := range args {
		reflectArgs = append(reflectArgs, reflect.ValueOf(arg))
	}

	// Call the method
	results := method.Call(reflectArgs)
	return results, nil
}

// Configuration loader using reflection
func loadConfigFromTags(v interface{}) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	fmt.Println("Loading config from struct tags...")

	// Handle nested structs
	loadStructConfig(val, typ, "")
}

func loadStructConfig(val reflect.Value, typ reflect.Type, prefix string) {
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		configTag := fieldType.Tag.Get("config")
		if configTag == "" {
			configTag = fieldType.Name
		}

		if prefix != "" {
			configTag = prefix + "." + configTag
		}

		if field.Kind() == reflect.Struct {
			loadStructConfig(field, field.Type(), configTag)
		} else {
			fmt.Printf("%s: %v\n", configTag, field.Interface())
		}
	}
}

// Generic operations using reflection
func processGenericValue(v interface{}) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	fmt.Printf("Processing %s: %v\n", typ, v)

	switch val.Kind() {
	case reflect.Slice:
		fmt.Printf("Type: %s\n", typ)
		fmt.Printf("Length: %d\n", val.Len())

		// Sum if it's a numeric slice
		if val.Len() > 0 && val.Index(0).Kind() == reflect.Int {
			sum := 0
			for i := 0; i < val.Len(); i++ {
				sum += int(val.Index(i).Int())
			}
			fmt.Printf("Sum: %d\n", sum)
		}

	case reflect.Map:
		fmt.Printf("Type: %s\n", typ)

		// Extract keys and values
		var keys, values []string
		for _, key := range val.MapKeys() {
			keys = append(keys, fmt.Sprintf("%v", key.Interface()))
			values = append(values, fmt.Sprintf("%v", val.MapIndex(key).Interface()))
		}
		fmt.Printf("Keys: %v\n", keys)
		fmt.Printf("Values: %v\n", values)

	default:
		fmt.Printf("Type: %s\n", typ)
		fmt.Printf("Value: %v\n", val.Interface())
	}
}

// Object serializer using reflection
func serializeObject(v interface{}) (string, error) {
	val := reflect.ValueOf(v)

	// Handle pointers
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Create a map to hold the serialized data
	result := make(map[string]interface{})

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)

		// Get JSON tag
		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = fieldType.Name
		}

		// Skip fields with json:"-"
		if jsonTag == "-" {
			continue
		}

		result[jsonTag] = field.Interface()
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// Type conversion using reflection
func convertType(value interface{}, targetType reflect.Type) (interface{}, error) {
	val := reflect.ValueOf(value)

	// Try to convert
	if val.CanConvert(targetType) {
		return val.Convert(targetType).Interface(), nil
	}

	// Handle specific conversions
	switch targetType.Kind() {
	case reflect.String:
		return fmt.Sprintf("%v", value), nil
	case reflect.Int:
		switch v := value.(type) {
		case string:
			return strconv.Atoi(v)
		case float64:
			return int(v), nil
		case bool:
			if v {
				return 1, nil
			}
			return 0, nil
		}
	case reflect.Float64:
		switch v := value.(type) {
		case string:
			return strconv.ParseFloat(v, 64)
		case int:
			return float64(v), nil
		}
	case reflect.Bool:
		switch v := value.(type) {
		case string:
			return strconv.ParseBool(v)
		case int:
			return v != 0, nil
		}
	}

	return nil, fmt.Errorf("cannot convert %v to %s", value, targetType)
}

func main() {
	fmt.Println("Exercise 12: Reflection")
	fmt.Println("=======================")

	// 1. Struct inspection
	fmt.Println("Struct Inspection:")
	person := Person{
		Name:   "John Doe",
		Age:    30,
		Email:  "john@example.com",
		Active: true,
	}

	inspectStruct(person)

	// 2. JSON validation
	fmt.Println("JSON Validation:")
	err := validateStruct(person)
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
	}
	fmt.Println()

	// 3. Dynamic method calls
	fmt.Println("Dynamic Method Calls:")

	// Call GetFullName method
	results, err := callMethod(person, "GetFullName")
	if err != nil {
		fmt.Printf("Error calling method: %v\n", err)
	} else {
		fmt.Printf("Calling method: GetFullName\n")
		fmt.Printf("Result: %s\n", results[0].String())
	}

	// Call IsAdult method
	results, err = callMethod(person, "IsAdult")
	if err != nil {
		fmt.Printf("Error calling method: %v\n", err)
	} else {
		fmt.Printf("Calling method: IsAdult\n")
		fmt.Printf("Result: %t\n", results[0].Bool())
	}
	fmt.Println()

	// 4. Configuration loading
	fmt.Println("Configuration Loading:")
	config := Config{}
	config.Server.Host = "localhost"
	config.Server.Port = 8080
	config.Database.Host = "db.example.com"
	config.Database.Port = 5432
	config.Database.Name = "myapp"

	loadConfigFromTags(&config)
	fmt.Println()

	// 5. Generic operations
	fmt.Println("Generic Operations:")

	// Process slice
	numbers := []int{1, 2, 3, 4, 5}
	processGenericValue(numbers)

	// Process map
	fruits := map[string]string{"apple": "red", "banana": "yellow"}
	processGenericValue(fruits)
	fmt.Println()

	// 6. Object serialization
	fmt.Println("Object Serialization:")
	jsonStr, err := serializeObject(person)
	if err != nil {
		fmt.Printf("Error serializing: %v\n", err)
	} else {
		fmt.Println("Serializing Person struct...")
		fmt.Println(jsonStr)
	}
	fmt.Println()

	// 7. Type conversion
	fmt.Println("Type Conversion:")

	// String to int
	if result, err := convertType("42", reflect.TypeOf(0)); err == nil {
		fmt.Printf("Converting string \"42\" to int: %v\n", result)
	}

	// Float to string
	if result, err := convertType(3.14, reflect.TypeOf("")); err == nil {
		fmt.Printf("Converting float 3.14 to string: %v\n", result)
	}

	// Bool to string
	if result, err := convertType(true, reflect.TypeOf("")); err == nil {
		fmt.Printf("Converting bool true to string: %v\n", result)
	}

	fmt.Println("\nExercise completed!")
}
