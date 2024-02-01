package main

import "fmt"

func main() {
	/*
	   mengambil sub string dengan tehnik array
	   dan kemudian mengkonversinya menjadi string,karena ketika di ambil akan berupa kode ascii jadi harus\
	   di konversi agar bisa di tampilkan sebagai huruf di terminal ketika di eksekusi
	*/
	nama := "dragon"
	sub := nama[2]
	sstring := string(sub)
	fmt.Println(nama)
	fmt.Println(sstring)
}
