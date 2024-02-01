package main

import "fmt"

/*biasanya kita memberi nama untuk sebuah function, tapi kita bisa juga
membuat function tanpa nama */
type f func(string) bool

func halo(nama string, spamm f) {
	if spamm(nama) { //jika hasil filternya true makan akan di blok
		fmt.Println("menggunakan kata jahat")
	} else {
		fmt.Println("halo ", nama)
	}
}

func main() {

	//contoh penggunaan anonymous functio
	bl := func(nama string) bool {
		return nama == "anjing"
	}

	halo("naga", bl)
	halo("anjing", func(nama string) bool {
		return nama == "anjing"
	})

	squareOf2 := func() int {
		return 2 * 2
	}()
	fmt.Println(squareOf2) // jka tidak menggunakan iifi squareOf2 akan bertipe fungsi

}
