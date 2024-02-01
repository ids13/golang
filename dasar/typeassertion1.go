package main

import "fmt"

func random() interface{} {
	return "halo"
}

func main() {
	data := random()
	resultdata := data.(string) //merubah tipe data / type assertion
	/* karena return dari fungsi random berupa text,maka bisa di rubah jadi string
	jika salah menentukan tipe data akan terjadi panic */
	fmt.Printf("tipe = %T\n isi = %v", resultdata, resultdata)
}
