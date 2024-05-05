package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Contoh penggunaan Read
	buf1 := bytes.NewBuffer([]byte("Hello, World!"))
	data1 := make([]byte, 5)
	n1, err := buf1.Read(data1)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Printf("Data read using Read: %s\n", data1[:n1])

	// Contoh penggunaan ReadByte
	buf2 := bytes.NewBuffer([]byte("Golang"))
	b2, err := buf2.ReadByte()
	if err != nil {
		fmt.Println("Error reading byte:", err)
		return
	}
	fmt.Printf("Byte read using ReadByte: %c\n", b2)

	// Contoh penggunaan ReadBytes
	buf3 := bytes.NewBuffer([]byte("is amazing!"))
	data3, err := buf3.ReadBytes(' ')
	if err != nil {
		fmt.Println("Error reading bytes:", err)
		return
	}
	fmt.Printf("Bytes read using ReadBytes: %s\n", data3)

	// Contoh penggunaan ReadFrom
	buf4 := bytes.NewBuffer([]byte("This is a test"))
	buf5 := bytes.NewBuffer(nil)
	n4, err := buf5.ReadFrom(buf4)
	if err != nil {
		fmt.Println("Error reading from buffer:", err)
		return
	}
	fmt.Printf("Bytes read from buffer using ReadFrom: %d\n", n4)

	// Contoh penggunaan Write
	var buf6 bytes.Buffer
	n6, err := buf6.Write([]byte("Hello"))
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Printf("Bytes written using Write: %d\n", n6)

	// Contoh penggunaan WriteByte
	var buf7 bytes.Buffer
	err = buf7.WriteByte('!')
	if err != nil {
		fmt.Println("Error writing byte:", err)
		return
	}
	fmt.Printf("Byte written using WriteByte\n")
}
