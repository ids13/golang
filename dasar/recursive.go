package main

import "fmt"

func fr(angka int) int {
	if angka == 1 {
		return 1
	} else {
		return angka * fr(angka-1) //memanggil fungsi dirinya sendiri
	}
}

func main() {
	fmt.Println(fr(5))
}
