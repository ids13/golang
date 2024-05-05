package main

import "fmt"

//alias untuk tipe data dasar (bisa dasar,map,slice dll)
type angka int

//alias untuk fungsi
type perbandingan func(a, b int) bool

//alias untuk struct
type orang struct {
	nama string
	umur int
}

// () di belakang nama fungsi menujukan receiver,membuat fungsi ini terikat dengan receiver,jadi seolah olah
//  receiver tersebut memiliki method
func (a angka) fangka(b int) string {
	return "angka dari alias a = " + string(a) + " dan angka yang di bypass ke fungsi angka = " + string(b)
}
func (a orang) forang() {
	fmt.Println(a.nama)
	fmt.Println(a.umur)
}
func (a perbandingan) fperbandingan(x, y int) bool {
	return a(x, y)
}
func kecil(a, b int) bool {
	return a < b
}
func besar(a, b int) bool {
	return a > b
}

func main() {
	var a angka = 1
	fmt.Println(a.fangka(2))
	x := 30
	y := 20
	//karena perbadingan adalah alias untuk fungsi tertentu
	//kita mengeset fungsi tersebut dengan fungsi lain yang memiliki format yang sama dengan fungsi yg ada di alias
	//kemudia memanggil method fperbadingan yang sudah terikat dengan alis perbandingan
	if perbandingan(kecil).fperbandingan(x, y) {
		fmt.Println("pertama lebih kecil dari kedua")
	} else if perbandingan(besar).fperbandingan(x, y) {
		fmt.Println("pertama lebih besar dari kedua")
	}
	orang{"andir", 20}.forang()
}
