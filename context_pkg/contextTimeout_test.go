package contextpkg

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func CounterTO(ctx context.Context, w *sync.WaitGroup) chan int {
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
				time.Sleep(1 * time.Second)
				counter++
			}
		}

	}()
	return dst
}

func TestCounterTO(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())

	var w sync.WaitGroup

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 3*time.Second)
	defer cancel()
	/*
		timeout adalah kejadian pembatalan ketika sudah melewati batas waktu
		bisa jadi prosesnya tidak melewati batas waktunya
		jadi cancel harus tetap di panggil,agar tidak ada proses leak
		jadi harus tetap menggunakan cance(), dalam kasus ini "defer cancel()"
	*/
	w.Add(1)
	destination := CounterTO(ctx, &w)
	fmt.Println(runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("counter ", n)
		if n == 10 {
			break
		}
	}

	w.Wait()
	fmt.Println(runtime.NumGoroutine())

}
