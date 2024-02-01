package main

import "fmt"

func halo() (string, string) {
	return "naga", "langit"
}

func main() {
	awal, akhir := halo() //di gunakan untuk menangkap hasil return secara berurutan
	fmt.Println(awal, akhir)
	/*multi return harus di tangkap semua hasil returnnya,jika ada yang ingin di hiraukan
	hasil returnya gunakan tanda garis bawah  */
	x, _ := halo()
	fmt.Println(x)
}
