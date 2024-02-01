package main

import (
	"fmt"
)

/*
	recover adalah fungsi yang di gunakan untuk menangkap data panic

dengan recover panic akan terhenti,dan program akan berjalan lagi
*/
func end() {
	message := recover()
	if message != nil {
		fmt.Println("pesan error:", message)
	}
	fmt.Println("aplikasi selesai")
}
func run(status bool) {
	defer end()
	if status {
		panic("aplikasi error") //ketika panic tereksekusi,kode di bawahnya tidak akan di eksekusi
	}
	fmt.Println("aplikasi berjalan")

}

func main() {
	run(true)
	fmt.Println("bukti aplikasi masih berjalan setelah di recover")
}
