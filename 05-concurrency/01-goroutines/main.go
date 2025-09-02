package main

import (
	"fmt"
)
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Msg is a typed envelope that flows through the pipeline.
type Msg[T any] struct {
	ID   int           // correlation / trace id
	Data T             // payload
	Err  error         // optional error; stages should pass through if set
	Meta map[string]any
}

// gen is a source stage.
func gen[T any](ctx context.Context, xs ...Msg[T]) <-chan Msg[T] {
	out := make(chan Msg[T])
	go func() {
		defer close(out)
		for _, x := range xs {
			select {
			case <-ctx.Done():
				return
			case out <- x:
			}
		}
	}()
	return out
}

// mapStage applies f to each message, keeping ID/Meta and propagating Errs.
func mapStage[I, O any](ctx context.Context, in <-chan Msg[I], f func(Msg[I]) Msg[O]) <-chan Msg[O] {
	out := make(chan Msg[O])
	go func() {
		defer close(out)
		for m := range in {
			if m.Err != nil {
				// propagate error without touching Data
				out <- Msg[O]{ID: m.ID, Err: m.Err, Meta: m.Meta}
				continue
			}
			select {
			case <-ctx.Done():
				return
			case out <- f(m):
			}
		}
	}()
	return out
}

// filterStage drops messages that don't satisfy pred.
func filterStage[T any](ctx context.Context, in <-chan Msg[T], pred func(Msg[T]) bool) <-chan Msg[T] {
	out := make(chan Msg[T])
	go func() {
		defer close(out)
		for m := range in {
			if m.Err != nil {
				out <- m
				continue
			}
			if pred(m) {
				select {
				case <-ctx.Done():
					return
				case out <- m:
				}
			}
		}
	}()
	return out
}

// fanout worker pool using structs keeps concurrency obvious.
func workerPool[I, O any](ctx context.Context, in <-chan Msg[I], n int, f func(context.Context, Msg[I]) Msg[O]) <-chan Msg[O] {
	out := make(chan Msg[O])
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for m := range in {
				if m.Err != nil {
					select {
					case <-ctx.Done():
						return
					case out <- Msg[O]{ID: m.ID, Err: m.Err, Meta: m.Meta}:
					}
					continue
				}
				res := f(ctx, m)
				select {
				case <-ctx.Done():
					return
				case out <- res:
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Source data
	src := []Msg[int]{
		{ID: 1, Data: 2}, {ID: 2, Data: 3}, {ID: 3, Data: 4}, {ID: 4, Data: 5},
	}

	// Build a readable pipeline with named stages.
	in := gen(ctx, src...)
	onlyEven := filterStage(ctx, in, func(m Msg[int]) bool { return m.Data%2 == 0 })
	squared := mapStage(ctx, onlyEven, func(m Msg[int]) Msg[int] {
		m.Data = m.Data * m.Data
		return Msg[int]{ID: m.ID, Data: m.Data, Meta: m.Meta}
	})
	// Simulate an I/O bound sink with a worker pool.
	written := workerPool(ctx, squared, 3, func(ctx context.Context, m Msg[int]) Msg[int] {
		time.Sleep(150 * time.Millisecond) // pretend to write to DB
		// Could set m.Meta["db"] = "ok" here
		return m
	})

	for m := range written {
		if m.Err != nil {
			fmt.Println("error:", m.ID, m.Err)
			continue
		}
		fmt.Printf("ID=%d out=%d\n", m.ID, m.Data)
	}
}

func main() {
	// TODO: Implement the goroutines exercise
	// 1. Create multiple goroutines that print messages
	// 2. Implement concurrent counter with mutex protection
	// 3. Create worker pool pattern
	// 4. Demonstrate race conditions and fixes
	// 5. Use channels for communication

	fmt.Println("Exercise 6: Goroutines")
	fmt.Println("======================")

	// Your implementation goes here
}
