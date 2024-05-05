package main

import (
	"fmt"
	"sync"
)

func main() {
	var m = make(map[int]int)
	var mu sync.Mutex // Mutex untuk mengamankan akses ke map
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		mu.Lock() // Mengunci map sebelum modifikasi
		for i := 0; i < 10; i++ {
			m[i] = i 
		}
		mu.Unlock() // Membuka kunci map setelah selesai modifikasi
	}()

	// Menunggu goroutine selesai menulis
	wg.Wait()

	for i := 0; i < 10; i++ {
		fmt.Println("Value for key", i, "is", m[i])
	}
}
