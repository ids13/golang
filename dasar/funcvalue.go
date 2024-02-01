package main

import "fmt"

/* function juga merupakan tipe data ,jadi bisa di simpan kedalam variable */

func halo(par string) string {
	return "halo " + par
}
func main() {
	x := halo //untuk memasukan fungsi ke dalam variable,masukan namanya saja,jagan ada tanda kurungnya
	fmt.Println(x("naga"))
}
