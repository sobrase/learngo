package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function for goroutines
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d: Starting...\n", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Worker %d: Processing...\n", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Worker %d: Completed\n", id)
}

// Counter with mutex protection
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *SafeCounter) Decrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count--
}

func (c *SafeCounter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// Unsafe counter for race condition demonstration
type UnsafeCounter struct {
	count int
}

func (c *UnsafeCounter) Increment() {
	c.count++
}

func (c *UnsafeCounter) GetValue() int {
	return c.count
}

// Worker pool implementation
func workerPool(workerID int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing task %d\n", workerID, job)
		time.Sleep(50 * time.Millisecond) // Simulate work
		results <- job * 2                // Process the job
	}
}

// Channel communication example
func sender(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- "Hello from goroutine"
	close(ch)
}

func receiver(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := <-ch
	fmt.Printf("Received: %s\n", msg)
}

func main() {
	fmt.Println("Exercise 6: Goroutines")
	fmt.Println("======================")

	// 1. Basic goroutines with WaitGroup
	fmt.Println("Goroutine Messages:")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println()

	// 2. Concurrent counter with mutex protection
	fmt.Println("Concurrent Counter:")
	safeCounter := &SafeCounter{}

	// Increment counter 1000 times concurrently
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			safeCounter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Initial value: 0\n")
	fmt.Printf("After 1000 increments: %d\n", safeCounter.GetValue())

	// Decrement counter 500 times concurrently
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			safeCounter.Decrement()
		}()
	}

	wg.Wait()
	fmt.Printf("After 500 decrements: %d\n", safeCounter.GetValue())
	fmt.Println()

	// 3. Worker pool pattern
	fmt.Println("Worker Pool:")
	numJobs := 5
	numWorkers := 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go workerPool(i, jobs, results, &wg)
	}

	// Send jobs
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	// Collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print results
	var resultSlice []int
	for result := range results {
		resultSlice = append(resultSlice, result)
	}

	fmt.Printf("All tasks completed\n")
	fmt.Println()

	// 4. Race condition demonstration
	fmt.Println("Race Condition Demo:")

	// Unsafe counter (may show race condition)
	unsafeCounter := &UnsafeCounter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			unsafeCounter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Without mutex (may show race condition):\n")
	fmt.Printf("Counter value: %d (should be 1000)\n", unsafeCounter.GetValue())
	fmt.Println()

	// Safe counter with mutex
	safeCounter2 := &SafeCounter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			safeCounter2.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("With mutex protection:\n")
	fmt.Printf("Counter value: %d\n", safeCounter2.GetValue())
	fmt.Println()

	// 5. Channel communication
	fmt.Println("Channel Communication:")
	ch := make(chan string)

	wg.Add(2)
	go sender(ch, &wg)
	go receiver(ch, &wg)

	wg.Wait()
	fmt.Println()

	// 6. Timeout with channels
	fmt.Println("Timeout Example:")
	timeoutCh := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		timeoutCh <- "Slow operation completed"
	}()

	select {
	case result := <-timeoutCh:
		fmt.Printf("Received: %s\n", result)
	case <-time.After(1 * time.Second):
		fmt.Println("Operation timed out")
	}
	fmt.Println()

	// 7. Buffered channels
	fmt.Println("Buffered Channel Demo:")
	bufferedCh := make(chan int, 3)

	// Send values to buffered channel
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3

	// Receive values
	fmt.Printf("Received: %d\n", <-bufferedCh)
	fmt.Printf("Received: %d\n", <-bufferedCh)
	fmt.Printf("Received: %d\n", <-bufferedCh)
	fmt.Println()

	// 8. Select statement with multiple channels
	fmt.Println("Select Statement Demo:")
	ch1 := make(chan string)
	ch2 := make(chan int)

	go func() {
		ch1 <- "Hello"
		time.Sleep(100 * time.Millisecond)
		ch1 <- "World"
	}()

	go func() {
		ch2 <- 42
		time.Sleep(50 * time.Millisecond)
		ch2 <- 100
	}()

	// Receive from multiple channels
	for i := 0; i < 4; i++ {
		select {
		case msg := <-ch1:
			fmt.Printf("Received from ch1: %s\n", msg)
		case num := <-ch2:
			fmt.Printf("Received from ch2: %d\n", num)
		case <-time.After(200 * time.Millisecond):
			fmt.Println("Timeout occurred")
		}
	}
}
