package contextpkg

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func CounterCancel(ctx context.Context, w *sync.WaitGroup) chan int {
	dst := make(chan int)
	go func() {
		defer close(dst)
		defer w.Done()
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return

			default:
				dst <- counter
				counter++
			}
		}

	}()
	return dst
}

func TestCounterCancel(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())

	var w sync.WaitGroup

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)
	w.Add(1)
	destination := CounterCancel(ctx, &w)

	for n := range destination {
		fmt.Println("counter ", n)
		if n == 10 {
			break
		}
	}

	cancel()
	w.Wait()
	fmt.Println(runtime.NumGoroutine())

}
