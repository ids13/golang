package main

import "fmt"

/* walaupun method akan menempel di struct datanya di akses bypass by value
sangat di rekomendasikan menggunakan pointer di method struct agar tidak boros memory */

type Man struct {
	name string
}

func (man *Man) nikah() { //penggunaan pointer
	man.name = "mr." + man.name
}

func main() {
	v := Man{"naga"} //menyimpan pointer struct man
	v.nikah()
	fmt.Println(v.name)
}
