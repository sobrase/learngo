# Go Learning Curriculum

This repository contains a comprehensive set of Go exercises designed to progressively teach you Go programming concepts and standard library usage.

## Structure

The exercises are organized by concept and difficulty level:

### Level 1: Fundamentals
- **01-basics** - Variables, types, functions, control structures
- **02-slices-maps** - Working with slices and maps
- **03-structs-methods** - Structs, methods, and interfaces
- **04-error-handling** - Error handling patterns

### Level 2: Intermediate
- **05-concurrency** - Goroutines and channels
- **06-io-operations** - File I/O and text processing
- **07-json-xml** - Data serialization
- **08-http-server** - HTTP server and client

### Level 3: Advanced
- **09-database** - Database operations with SQL
- **10-testing** - Unit testing and benchmarking
- **11-reflection** - Reflection and metaprogramming
- **12-web-framework** - Building a web application

## How to Use

1. Start with the first exercise in each folder
2. Read the problem description in `README.md`
3. Implement your solution in `main.go`
4. Run tests with `go test`
5. Check the solution in `solution.go` if needed

## Prerequisites

- Go 1.21 or later
- Basic programming knowledge

## Getting Started

```bash
# Navigate to any exercise folder
cd 01-basics/01-variables

# Run the exercise
go run main.go

# Run tests
go test

# Check the solution
go run solution.go
```

## Standard Library Coverage

This curriculum covers the most important Go standard library packages:

- `fmt` - Formatting and printing
- `strings` - String manipulation
- `strconv` - String conversions
- `math` - Mathematical functions
- `time` - Time and date operations
- `os` - Operating system interface
- `io` - I/O primitives
- `bufio` - Buffered I/O
- `encoding/json` - JSON encoding/decoding
- `encoding/xml` - XML encoding/decoding
- `net/http` - HTTP client and server
- `database/sql` - SQL database interface
- `testing` - Testing framework
- `reflect` - Reflection
- `sync` - Synchronization primitives
- `context` - Context management
- `log` - Logging
- `regexp` - Regular expressions
- `sort` - Sorting algorithms
- `container/heap` - Heap data structure

Happy coding! 🚀
