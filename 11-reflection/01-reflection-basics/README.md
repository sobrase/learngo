# Exercise 12: Reflection

## Objective
Learn about Go's `reflect` package and metaprogramming capabilities.

## Problem
Create a program that demonstrates reflection for type inspection, struct field access, and dynamic function calls.

## Requirements
1. Use `reflect` package for type inspection:
   - `reflect.TypeOf()` and `reflect.ValueOf()`
   - Type and value information
2. Work with struct fields using reflection:
   - Field iteration and access
   - Tag reading and processing
3. Use reflection for function calls:
   - Dynamic method invocation
   - Parameter validation
4. Implement generic operations using reflection
5. Handle different data types dynamically

## Tasks
1. Create a struct inspector that prints field information
2. Implement a generic JSON validator using reflection
3. Create a function that dynamically calls methods
4. Build a configuration loader using struct tags
5. Implement a simple object serializer

## Expected Output
```
Reflection Demo:

Struct Inspection:
Type: main.Person
Kind: struct
Num fields: 4

Field 0: Name (string)
  Tag: json:"name" validate:"required"
  Value: John Doe
  Can set: true

Field 1: Age (int)
  Tag: json:"age" validate:"min:18"
  Value: 30
  Can set: true

Field 2: Email (string)
  Tag: json:"email" validate:"email"
  Value: john@example.com
  Can set: true

Field 3: Active (bool)
  Tag: json:"active"
  Value: true
  Can set: true

JSON Validation:
Validating Person struct...
✓ Name: required field present
✓ Age: 30 >= 18 (min requirement met)
✓ Email: valid email format
✓ Active: boolean field valid

Dynamic Method Calls:
Calling method: GetFullName
Result: John Doe (30)

Calling method: IsAdult
Result: true

Configuration Loading:
Loading config from struct tags...
Server:
  Host: localhost
  Port: 8080
Database:
  Host: db.example.com
  Port: 5432
  Name: myapp

Generic Operations:
Processing slice: [1, 2, 3, 4, 5]
Type: []int
Length: 5
Sum: 15

Processing map: map[apple:red banana:yellow]
Type: map[string]string
Keys: [apple, banana]
Values: [red, yellow]

Object Serialization:
Serializing Person struct...
{
  "name": "John Doe",
  "age": 30,
  "email": "john@example.com",
  "active": true
}

Type Conversion:
Converting string "42" to int: 42
Converting float 3.14 to string: 3.14
Converting bool true to string: true
```

## Hints
- Use `reflect.TypeOf()` to get type information
- Use `reflect.ValueOf()` to get value information
- Use `reflect.Value.Elem()` to get the value pointed to by a pointer
- Use `reflect.Value.CanSet()` to check if a value can be modified
- Use `reflect.Value.Field()` to access struct fields
- Use `reflect.Value.Method()` to access methods
- Use `reflect.Value.Call()` to invoke functions
- Remember that reflection is slower than direct access
- Use `reflect.Kind` to check the underlying type
- Handle panics when using reflection
