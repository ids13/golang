package main

import "fmt"

func main() {
	var nilai interface{}
	nilai = 3
	// nilai.(type) : format (type) ini hanya berlaku di switch
	switch v := nilai.(type) {
	case string:
		fmt.Println("Ini adalah sebuah string:", v)
	case int:
		fmt.Println("Ini adalah sebuah angka bulat:", v)
	case float64:
		fmt.Println("Ini adalah sebuah angka desimal:", v)
	default:
		fmt.Println("Tipe nilai tidak dikenali.")
	}
}
