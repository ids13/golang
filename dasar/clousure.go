package main

import "fmt"

/* clousere adalah kemampuan function untuk berinteraksi dengan data-data di sekitarnya dengan lingkup scope
yang sama */
func main() {
	c := 0
	inc := func() {
		fmt.Println("increment")
		c++ //ketika c di increment c di luar fungsi (masih dalam scope funsi main) akan terpengaruh
	}
	/* data yang di atas bisa di akses dari dalam fungsi
	jika sebaliknya maka tidak bisa.maka dari itu harus berhati hati dalam mekases data di dalam fungsi
	jika ingin tidak berubah.harus melakukan pendeklarasian ulang,jadi meskipun namanya sama,akan di anggap
	berbeda oleh golang */
	inc()
	inc()
	fmt.Println(c)
}
