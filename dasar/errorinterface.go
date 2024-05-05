package main

import (
	"errors"
	"fmt"
)

// Fungsi untuk memeriksa apakah sebuah kata memiliki panjang yang mencukupi
func cekPanjangKata(kata string) error {
	if len(kata) < 5 {
		return errors.New("panjang kata terlalu pendek")
	}
	return nil
}

func main() {
	// Contoh 1: Sukses
	var err error
	err = cekPanjangKata("contoh")
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Panjang kata mencukupi.")
	}

	// Contoh 2: Gagal (kata terlalu pendek)
	err = cekPanjangKata("go")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Panjang kata mencukupi.")

	}
}
