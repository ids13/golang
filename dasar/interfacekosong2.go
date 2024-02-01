package main

import "fmt"

// Fungsi yang mengembalikan nilai bertipe interface kosong
//interface kosong bisa juga dengan mengtikan langsung "interface{}"
func processData(input interface{}) interface{} {
	// Lakukan operasi atau manipulasi data sesuai kebutuhan
	// Pada contoh ini, kita hanya mengembalikan data tanpa perubahan
	return input
}

func main() {
	// Contoh penggunaan interface kosong dengan berbagai tipe data
	var stringData interface{} = "Hello, Golang!"
	var intData interface{} = 42
	var floatData interface{} = 3.14
	var boolData interface{} = true

	// Memanggil fungsi processData dengan berbagai tipe data
	resultString := processData(stringData)
	resultInt := processData(intData)
	resultFloat := processData(floatData)
	resultBool := processData(boolData)

	// Menampilkan informasi hasil pemrosesan
	fmt.Printf("Hasil String: %v, Tipe: %T\n", resultString, resultString)
	fmt.Printf("Hasil Int: %v, Tipe: %T\n", resultInt, resultInt)
	fmt.Printf("Hasil Float: %v, Tipe: %T\n", resultFloat, resultFloat)
	fmt.Printf("Hasil Bool: %v, Tipe: %T\n", resultBool, resultBool)
}
