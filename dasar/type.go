package main

import "fmt"

/*
di gunakan untuk membuat alias tipe data
*/
func main() {
	type nktp string
	type status bool

	var nama nktp = "dragon"
	var kuat status = true

	fmt.Println(nama)
	fmt.Println(kuat)

}
