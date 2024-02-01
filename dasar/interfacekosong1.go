package main

import "fmt"

/* interface kosong adalah interface yang tidak memiliki deklarasi fungsi satupun,hal ini
membuat secara otomatis semua tipe data menjadi implementasinya */
// Buatkan interface kosong
type EmptyInterface interface{}

// Fungsi yang menerima parameter dengan tipe interface kosong
func printData(data EmptyInterface) {
	fmt.Println("Data:", data)
}

func main() {
	// Contoh penggunaan interface kosong dengan tipe data yang berbeda
	var stringData EmptyInterface = "Hello, Golang!"
	var intData EmptyInterface = 42
	var floatData EmptyInterface = 3.14

	// Panggil fungsi printData dengan berbagai tipe data
	printData(stringData)
	printData(intData)
	printData(floatData)
}
