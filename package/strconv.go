package main

import (
	"fmt"
	"strconv"
)

/*
	strconv digunakna untuk conversi tipe data dari atau ke string

parse : dari string
format : ke string
*/
func main() {
	// merubah string true menjadi boolean dan menampungnya ke dalam variable hand, jika di cek tipe data hand akan jadi bolean dengan nilai true
	hand, err := strconv.ParseBool("true")
	fmt.Println("hand ", hand, " error ", err)
	// int ke format string, 10 untuk oktal,16 untuk hexa,
	fmt.Println(strconv.FormatInt(100000, 16))
	// convert ke int , tinggal abaikan nilai err nya
	fmt.Println(strconv.Atoi("1000000"))

}
