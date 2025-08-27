# Exercise 7: Channels

## Objective
Learn about Go channels, buffered channels, and channel patterns.

## Problem
Create a program that demonstrates channel operations, buffered channels, and common channel patterns.

## Requirements
1. Create and use unbuffered channels
2. Work with buffered channels
3. Use `select` statements for non-blocking operations
4. Implement channel patterns:
   - Producer-consumer
   - Fan-out/Fan-in
   - Pipeline
5. Use `close()` and range over channels
6. Handle channel timeouts with `time.After()`

## Tasks
1. Create a producer-consumer pattern
2. Implement a pipeline for data processing
3. Use select statements for multiple channels
4. Create a fan-out/fan-in pattern
5. Demonstrate channel timeouts and non-blocking operations

## Expected Output
```
Producer-Consumer Pattern:
Produced: 1
Produced: 2
Produced: 3
Consumed: 1
Consumed: 2
Consumed: 3

Pipeline Processing:
Original: [1, 2, 3, 4, 5]
Squared: [1, 4, 9, 16, 25]
Doubled: [2, 8, 18, 32, 50]

Select Statement Demo:
Received from channel1: A
Received from channel2: 1
Received from channel1: B
Received from channel2: 2
Timeout occurred

Fan-Out/Fan-In Pattern:
Input: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
Worker 1 processing: 1
Worker 2 processing: 2
Worker 1 processing: 3
Worker 2 processing: 4
Worker 1 processing: 5
Worker 2 processing: 6
Worker 1 processing: 7
Worker 2 processing: 8
Worker 1 processing: 9
Worker 2 processing: 10
Results: [2, 4, 6, 8, 10, 12, 14, 16, 18, 20]

Buffered Channel Demo:
Sending to buffered channel...
Sent: 1
Sent: 2
Sent: 3
Received: 1
Received: 2
Received: 3

Non-blocking Operations:
Channel is ready: false
Default case executed
```

## Hints
- Use `make(chan T)` for unbuffered channels
- Use `make(chan T, size)` for buffered channels
- Use `close(channel)` to close channels
- Use `range channel` to iterate over channel values
- Use `select` for non-blocking channel operations
- Remember that sending on a closed channel causes panic
- Use `time.After()` for channel timeouts
- Channels are the primary way goroutines communicate
