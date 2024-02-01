package main

import "fmt"

/* membuat fungsi sebagai parameter
untuk contoh di bawah, "fungsi func(string) string" bisa di singkat menggunakan type
type di gunakan untuk membuat tipe data baru dari data yang sudah ada
contoh:

type f func(string) string
jadi di tulis

func halo(nama string, fungsi f)
*/

func halo(nama string, fungsi func(string) string) {
	fmt.Println("hallo ", fungsi(nama))
}
func spamm(kata string) string {
	if kata == "anjing" {
		return "..."
	} else {
		return kata
	}
}

func main() {
	halo("naga", spamm)
	halo("anjing", spamm)
}
