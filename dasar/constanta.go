package main

import "fmt"

func main() {
	/*
	   -inisialisasi constanta seperti variable perbedaannya
	   variable menggunakan "var" constanta menggunakan "const"
	   -harus langsung di isi nilainya ketika di deklarasikan
	   -jika tidak di gunakan tidak akan menimbulkan erro berbeda dengan variable
	   -tidab bisa menggunakan tanda ":=". jika di gunakan makan golang akan membuatnya menjadi variable
	*/
	const name1 string = "dragon"
	const name2 = "fly"
	fmt.Println(name1)
	fmt.Println(name2)

	const (
		data1 = "multi"
		data2 = 123
	)
	fmt.Println(data1)
	fmt.Println(data2)

}
