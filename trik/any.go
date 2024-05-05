package main

import "fmt"

// tanda kurung siku menunjukkan bahwa fungsi tersebut memiliki parameter generik
func printAnyType[T any](value T) {
	fmt.Println(value)
}

func main() {
	printAnyType("Hello, World!") // T akan diinferensikan sebagai string
	printAnyType(42)              // T akan diinferensikan sebagai int
	printAnyType(3.14)            // T akan diinferensikan sebagai float64
}

/*
tipe data any hanay tersedi adi versi golang versi 1.18 ke atas
Kekurangan dari interface{} adalah bahwa kita perlu melakukan asertasi tipe (type assertion) atau tipe switch untuk mendapatkan nilai asli dan menggunakan nilainya dengan aman
any pada fitur generik Go 1.18 memberikan kemampuan untuk membuat fungsi atau struktur data yang dapat bekerja dengan tipe data yang berubah-ubah, tetapi dengan keamanan tipe pada tingkat kompilasi. */
