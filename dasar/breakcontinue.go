package main

import "fmt"

/* break akan menghentikan looping
continue melakukan skip looping dan melanjukan ke looping berikutnya */

func main() {
	for i := 0; i < 10; i++ {
		if i == 6 {
			break //akan berhenti di angka lima,karena setelah ketemu angka enam akan langsung di break
		}
		fmt.Println("angka = ", i)
	}
	for i := 0; i < 10; i++ {
		if i == 4 {
			continue
		}
		fmt.Println("angka = ", i)
	}

}
