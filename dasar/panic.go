package main

import (
	"fmt"
)

/*
	panic adalah fungsi yang di gunakan untuk menghentikan program

ketika di panggil ,program akan terhenti namun defer akan tetap di jalankan
*/
func end() {
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
}
