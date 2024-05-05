package contextpkg

import (
	"fmt"
	"runtime"
	"testing"
)

func Counter() chan int {
	dst := make(chan int)
	go func() {
		defer close(dst)
		counter := 1
		for {
			dst <- counter
			counter++
		}
		/* 
		baris for tanpa ujung ini penyebab goroutine leak, karena datanya akan terus terusan di kirim ke channel
		*/
	}()
	return dst
}

func TestCounter(t *testing.T) {
	fmt.Println(runtime.NumGoroutine()) //jumlah sebelum di jalan kan (hasilnya 2)
 
	destination := Counter()
	for n := range destination {
		fmt.Println("counter ",n)
		if n == 10 {
			break
		}
	}
	fmt.Println(runtime.NumGoroutine()) //jumlah sesudah di jalan kan (hasilnya 3)
	/* 
	seharusnya sedah oprasi jumlah go routinya sama tapi ini berbeda. intinya goroutine nya jalan terus
	tanpa bisa di hentikan ini yang di namakan goroutine leak
	*/
}
