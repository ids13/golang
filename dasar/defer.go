package main

import "fmt"

/* defer adalah sebuah fungsi yang di eksekusi setelah fungsi selesai di eksekusi
fungsi defer akan di eksekusi ,meskipun terjadi error*/
func log() {
	fmt.Println("log di eksekusi")
}

func run(r int) {
	defer log() //akan di eksekusi meskipun terjadi error
	fmt.Println("run di eksekusi")
	r = 10 / r
	fmt.Println(r)
}

func main() {
	run(0) //karena nantinya di bagi nol akan di bagi terjadi error
}
