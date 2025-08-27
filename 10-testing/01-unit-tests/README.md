# Exercise 9: Unit Testing

## Objective
Learn about Go's testing framework, writing unit tests, and benchmarking.

## Problem
Create comprehensive unit tests for various functions and demonstrate testing best practices.

## Requirements
1. Write unit tests using the `testing` package
2. Use table-driven tests for multiple test cases
3. Write benchmarks using `testing.B`
4. Use test helpers and setup/teardown functions
5. Test error conditions and edge cases
6. Use `testing.T` for assertions and error reporting

## Tasks
1. Write tests for mathematical functions
2. Create table-driven tests for string operations
3. Write benchmarks for performance testing
4. Test struct methods and interfaces
5. Implement test helpers and utilities

## Expected Output
```
Running tests:
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
=== RUN   TestMultiply
--- PASS: TestMultiply (0.00s)
=== RUN   TestDivide
--- PASS: TestDivide (0.00s)
    divide_test.go:25: Testing division by zero
--- PASS: TestDivide (0.00s)
=== RUN   TestStringReverse
--- PASS: TestStringReverse (0.00s)
=== RUN   TestStringReverseTable
--- PASS: TestStringReverseTable (0.00s)
=== RUN   TestCalculator
--- PASS: TestCalculator (0.00s)
PASS
coverage: 85.7% of statements

Running benchmarks:
goos: windows
goarch: amd64
BenchmarkAdd-8           1000000000               0.0000001 ns/op
BenchmarkMultiply-8      1000000000               0.0000001 ns/op
BenchmarkStringReverse-8 10000000                150 ns/op
BenchmarkCalculator-8    10000000                200 ns/op
PASS

Test Results Summary:
✓ Add function: 5/5 test cases passed
✓ Multiply function: 3/3 test cases passed
✓ Divide function: 4/4 test cases passed (including error cases)
✓ String operations: 6/6 test cases passed
✓ Calculator struct: 8/8 test cases passed
```

## Hints
- Use `testing.T` for test functions
- Use `testing.B` for benchmark functions
- Use `t.Run()` for subtests
- Use `t.Helper()` for test helper functions
- Use `t.Fatal()` and `t.Error()` for test failures
- Use `testing.Short()` for quick tests
- Use `go test -v` for verbose output
- Use `go test -bench=.` for benchmarks
- Use `go test -cover` for coverage reports
