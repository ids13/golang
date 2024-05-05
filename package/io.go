package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	// Membuat buffer byte kosong
	var buffer bytes.Buffer
	buffer.WriteString("Hello, ")           // Menulis data ke buffer menggunakan WriteString
	buffer.WriteByte('W')                   // Menulis byte tunggal ke buffer menggunakan WriteByte
	data := []byte{'o', 'r', 'l', 'd', '!'} // Menambahkan byte menggunakan Write
	buffer.Write(data)

	readData := make([]byte, buffer.Len())
	n, _ := buffer.Read(readData) //Membaca data dari buffer menggunakan Read
	// Mencetak data yang telah dibaca dari buffer
	fmt.Printf("Data yang dibaca dari buffer: %s\n", readData[:n]) // Output: Hello, World!

	// Mengatur kembali posisi pembacaan ke awal buffer
	buffer.Reset()

	// Menulis data baru ke buffer
	buffer.WriteString("Selamat datang!")

	firstChar, _ := buffer.ReadByte()                           // Membaca satu karakter dari buffer menggunakan ReadByte
	fmt.Printf("Karakter pertama dari buffer: %c\n", firstChar) // Output: S

	str, _ := buffer.ReadString(' ')            // Membaca string dari buffer menggunakan ReadString
	fmt.Printf("String dari buffer: %s\n", str) // Output: elamat
	// Mencetak sisa konten dari buffer
	fmt.Printf("Sisa konten dari buffer: %s\n", buffer.String()) // Output: datang!
	//ketika menggunakan read, offset nya akan bergerak maju.

	var buff bytes.Buffer
	_, err := io.WriteString(&buff, "hallo dunia")
	if err != nil {
		println("error : ", err)
	}
	x, err := io.ReadAll(&buff)
	if err != nil {
		println("error : ", err)
	}
	println("data ada", len(x), "isinya adalah", string(x))
}
