package main

import "fmt"

func halo() (awal, tengah, akhir string) {
	awal = "naga"
	tengah = "langit"
	akhir = "bumi"
	return //cukup menggunakan "return" maka semua data akan di return(tidak perlu menyebutkan variable satu satu)
}

func main() {
	x, y, z := halo()
	fmt.Println(x, y, z)
}
