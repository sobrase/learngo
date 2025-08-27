# Exercise 6: Goroutines

## Objective
Learn about Go goroutines, basic concurrency, and the `sync` package.

## Problem
Create a program that demonstrates goroutines, synchronization, and concurrent operations.

## Requirements
1. Create and run goroutines using `go` keyword
2. Use `sync.WaitGroup` to wait for goroutines to complete
3. Use `sync.Mutex` for thread-safe operations
4. Demonstrate race conditions and their solutions
5. Use `time.Sleep()` and `time.After()` for timing

## Tasks
1. Create multiple goroutines that print messages
2. Implement a concurrent counter with mutex protection
3. Create a worker pool pattern
4. Demonstrate race conditions and fixes
5. Use channels for communication between goroutines

## Expected Output
```
Goroutine Messages:
Worker 1: Starting...
Worker 2: Starting...
Worker 3: Starting...
Worker 1: Processing...
Worker 2: Processing...
Worker 3: Processing...
Worker 1: Completed
Worker 2: Completed
Worker 3: Completed

Concurrent Counter:
Initial value: 0
After 1000 increments: 1000
After 500 decrements: 500

Worker Pool:
Worker 1 processing task 1
Worker 2 processing task 2
Worker 3 processing task 3
Worker 1 processing task 4
Worker 2 processing task 5
All tasks completed

Race Condition Demo:
Without mutex (may show race condition):
Counter value: 998 (should be 1000)

With mutex protection:
Counter value: 1000

Channel Communication:
Sending: Hello from goroutine
Received: Hello from goroutine
```

## Hints
- Use `go functionName()` to start a goroutine
- Use `sync.WaitGroup` to wait for goroutines to finish
- Use `sync.Mutex` to protect shared resources
- Use `defer wg.Done()` to ensure WaitGroup is decremented
- Remember that goroutines run concurrently, not necessarily in parallel
- Use channels for communication between goroutines
- Be careful with shared variables in concurrent code
