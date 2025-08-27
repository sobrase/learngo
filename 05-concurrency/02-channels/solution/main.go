package main

import (
	"fmt"
	"time"
)

// Producer-consumer pattern
func producer(ch chan<- int, items []int) {
	for _, item := range items {
		fmt.Printf("Produced: %d\n", item)
		ch <- item
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch <-chan int, done chan<- bool) {
	for item := range ch {
		fmt.Printf("Consumed: %d\n", item)
		time.Sleep(150 * time.Millisecond)
	}
	done <- true
}

// Pipeline pattern
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * 2
		}
	}()
	return out
}

// Fan-out/Fan-in pattern
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing: %d\n", id, job)
		time.Sleep(50 * time.Millisecond)
		results <- job * 2
	}
}

func main() {
	fmt.Println("Exercise 7: Channels")
	fmt.Println("====================")

	// 1. Producer-Consumer Pattern
	fmt.Println("Producer-Consumer Pattern:")
	items := []int{1, 2, 3}
	ch := make(chan int)
	done := make(chan bool)

	go producer(ch, items)
	go consumer(ch, done)

	<-done
	fmt.Println()

	// 2. Pipeline Processing
	fmt.Println("Pipeline Processing:")
	numbers := []int{1, 2, 3, 4, 5}

	// Create input channel
	input := make(chan int)
	go func() {
		defer close(input)
		for _, n := range numbers {
			input <- n
		}
	}()

	// Create pipeline: input -> square -> double
	squared := square(input)
	doubled := double(squared)

	// Collect results
	var results []int
	for result := range doubled {
		results = append(results, result)
	}

	fmt.Printf("Original: %v\n", numbers)

	// Calculate what the pipeline should produce
	var squaredResults []int
	for _, n := range numbers {
		squaredResults = append(squaredResults, n*n)
	}
	fmt.Printf("Squared: %v\n", squaredResults)
	fmt.Printf("Doubled: %v\n", results)
	fmt.Println()

	// 3. Select Statement Demo
	fmt.Println("Select Statement Demo:")
	ch1 := make(chan string)
	ch2 := make(chan int)

	go func() {
		ch1 <- "A"
		time.Sleep(100 * time.Millisecond)
		ch1 <- "B"
	}()

	go func() {
		ch2 <- 1
		time.Sleep(50 * time.Millisecond)
		ch2 <- 2
	}()

	// Receive from multiple channels
	for i := 0; i < 4; i++ {
		select {
		case msg := <-ch1:
			fmt.Printf("Received from channel1: %s\n", msg)
		case num := <-ch2:
			fmt.Printf("Received from channel2: %d\n", num)
		case <-time.After(300 * time.Millisecond):
			fmt.Println("Timeout occurred")
		}
	}
	fmt.Println()

	// 4. Fan-Out/Fan-In Pattern
	fmt.Println("Fan-Out/Fan-In Pattern:")
	input2 := make(chan int)
	output := make(chan int)

	// Start workers (fan-out)
	numWorkers := 2
	for i := 1; i <= numWorkers; i++ {
		go worker(i, input2, output)
	}

	// Send work (fan-out)
	go func() {
		defer close(input2)
		for i := 1; i <= 10; i++ {
			input2 <- i
		}
	}()

	// Collect results (fan-in)
	go func() {
		defer close(output)
	}()

	var finalResults []int
	for result := range output {
		finalResults = append(finalResults, result)
	}

	fmt.Printf("Input: %v\n", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Printf("Results: %v\n", finalResults)
	fmt.Println()

	// 5. Buffered Channel Demo
	fmt.Println("Buffered Channel Demo:")
	bufferedCh := make(chan int, 3)

	fmt.Println("Sending to buffered channel...")
	bufferedCh <- 1
	fmt.Println("Sent: 1")
	bufferedCh <- 2
	fmt.Println("Sent: 2")
	bufferedCh <- 3
	fmt.Println("Sent: 3")

	fmt.Printf("Received: %d\n", <-bufferedCh)
	fmt.Printf("Received: %d\n", <-bufferedCh)
	fmt.Printf("Received: %d\n", <-bufferedCh)
	fmt.Println()

	// 6. Non-blocking Operations
	fmt.Println("Non-blocking Operations:")
	nonBlockingCh := make(chan int)

	select {
	case value := <-nonBlockingCh:
		fmt.Printf("Channel is ready: %d\n", value)
	default:
		fmt.Println("Channel is ready: false")
		fmt.Println("Default case executed")
	}
	fmt.Println()

	// 7. Channel Timeouts
	fmt.Println("Channel Timeouts:")
	timeoutCh := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		timeoutCh <- "Slow operation completed"
	}()

	select {
	case result := <-timeoutCh:
		fmt.Printf("Operation completed: %s\n", result)
	case <-time.After(1 * time.Second):
		fmt.Println("Operation timed out")
	}
	fmt.Println()

	// 8. Channel Direction
	fmt.Println("Channel Direction Demo:")

	// Send-only channel
	sendCh := make(chan<- int)

	// Receive-only channel
	receiveCh := make(<-chan int)

	// Bidirectional channel
	bidirCh := make(chan int)

	// Demonstrate channel direction
	go func(ch chan<- int) {
		ch <- 42
	}(bidirCh)

	value := <-bidirCh
	fmt.Printf("Received value: %d\n", value)

	// Note: sendCh and receiveCh are for demonstration only
	// They can't be used in this example as they're unidirectional
	fmt.Println()

	// 9. Range over Channel
	fmt.Println("Range over Channel:")
	rangeCh := make(chan int)

	go func() {
		defer close(rangeCh)
		for i := 1; i <= 5; i++ {
			rangeCh <- i
			time.Sleep(50 * time.Millisecond)
		}
	}()

	fmt.Println("Values from channel:")
	for value := range rangeCh {
		fmt.Printf("  %d\n", value)
	}
	fmt.Println()

	// 10. Channel Capacity and Length
	fmt.Println("Channel Capacity and Length:")
	capCh := make(chan int, 5)

	fmt.Printf("Channel capacity: %d\n", cap(capCh))
	fmt.Printf("Channel length: %d\n", len(capCh))

	capCh <- 1
	capCh <- 2
	capCh <- 3

	fmt.Printf("After sending 3 values - length: %d\n", len(capCh))

	<-capCh
	fmt.Printf("After receiving 1 value - length: %d\n", len(capCh))
}
