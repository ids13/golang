package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Inisialisasi buffer byte
	buf := bytes.NewBufferString("Hello, World!")

	// Cetak isi buffer sebelum seek
	fmt.Println("Sebelum seek:", buf.String())

	// Lakukan seek ke posisi ke-7 (dalam byte)
	// Dalam hal ini, kita akan membaca 7 byte pertama dan buang
	// Ini akan mensimulasikan perpindahan posisi offset
	_, err := buf.Read(make([]byte, 7))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Tulis data baru setelah seek
	// Karena kita sudah melakukan seek, ini akan menambahkan data
	// setelah posisi offset yang baru
	_, err = buf.WriteString(" Gopher!")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Cetak isi buffer setelah seek dan penulisan
	fmt.Println("Setelah seek:", buf.String())
}
