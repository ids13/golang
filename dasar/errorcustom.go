package main

import "fmt"

// Custom error struct
type MyError struct {
	Message string
}

// Implementasi metode Error() untuk memenuhi interface error
func (e *MyError) Error() string {
	return e.Message
}

func main() {
	// Membuat instance custom error
	// "instance" mengacu pada penciptaan variabel dari suatu tipe data, baik itu tipe data
	err := &MyError{"Ini adalah custom error"}
	// Menangkap dan mencetak error
	fmt.Println("Error:", err)
}

/* kita  tidak perlu mengimpor package errors karena Anda tidak menggunakan errors.New()
kita hanya memanfaatkan kontrak interface error saja (sudah builtin seperti alias data bawaan lainnya)
dengan custom error kita bisa menkostumisasi data apa saja yang ingin di tampilkan*/
