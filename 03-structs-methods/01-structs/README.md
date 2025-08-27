# Exercise 5: Structs and Methods

## Objective
Learn about Go structs, methods, and object-oriented programming concepts.

## Problem
Create a program that demonstrates struct definition, methods, and interfaces.

## Requirements
1. Define structs with different field types:
   - Basic types (int, string, bool)
   - Slices and maps
   - Other structs (composition)
2. Create methods for structs:
   - Value receivers
   - Pointer receivers
   - Methods with different signatures
3. Use struct composition and embedding
4. Implement interfaces
5. Work with struct tags

## Tasks
1. Create a `Person` struct with methods
2. Create a `Rectangle` struct with area and perimeter methods
3. Implement a `BankAccount` struct with transaction methods
4. Create an interface for shapes
5. Use struct tags for JSON serialization

## Expected Output
```
Person Information:
Name: John Doe
Age: 30
Email: john@example.com
Full Info: John Doe (30) - john@example.com

Rectangle Operations:
Rectangle: Width=10, Height=5
Area: 50
Perimeter: 30
Is Square: false

Bank Account Operations:
Account: 12345
Balance: $1000.00
After deposit $500: $1500.00
After withdrawal $200: $1300.00
Transaction History:
  Deposit: +$500.00
  Withdrawal: -$200.00

Shape Interface:
Circle (radius=5): Area=78.54, Perimeter=31.42
Rectangle (10x5): Area=50.00, Perimeter=30.00

JSON Serialization:
{"name":"Jane Doe","age":25,"email":"jane@example.com","active":true}
```

## Hints
- Use `type StructName struct` to define structs
- Use `func (receiver Type) MethodName()` for methods
- Use pointer receivers `(receiver *Type)` when you need to modify the struct
- Use `interface{}` to define interfaces
- Use struct tags like `json:"field_name"` for serialization
- Remember that Go uses composition over inheritance
