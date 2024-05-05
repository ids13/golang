package main

import "fmt"

func main() {
	/*
		Slice adalah struktur data dinamis dalam bahasa pemrograman Go
		yang memungkinkan Anda untuk membuat dan bekerja dengan koleksi elemen.
		Slice dibangun di atas array dan memberikan cara yang lebih fleksibel untuk mengelola data.

				capacity di hitung dari elemen pertama yang di slice(paling kiri)
				len di hitung dari jumlah elemen yang di tampilkan dalam batasan

			a := [3]int{1,2,3} ini array
			a := [...]int{1,2,3} ini array
			a := []int{1,2,3} ini slice
	*/
	nama := []string{"apel", "mangga", "jeruk", "nangka"}
	fmt.Println(nama)
	fmt.Println(len(nama))
	fmt.Println(cap(nama))
	/*
	  contoh di bawah
	  1 adalah index batas bawahnya
	  3 adalah index batas atasnya, jadi akan menslice dari index ke 1 sampai index ke 2(dari 3-1)
	  len (hasil 2) menampilkan jumlah data yang di slice dalam batasan [1:3]
	  cap (hasil 3) kapasitas sisa dari slice asal(batas bawah). di hitung dari elemen pertama yang di slice(paling kiri).
	  rumusnya [cap asal ( dalam kasus ini 4) dikurangi index batas bawah(dalam kasus ini 1) maka hasilnya 3]
	*/
	anama := nama[1:3]
	fmt.Println(anama)
	fmt.Println(len(anama))
	fmt.Println(cap(anama))

	bnama := nama[2:]
	fmt.Println(bnama)
	fmt.Println(len(bnama))
	fmt.Println(cap(bnama))
	cnama := nama[:2]
	fmt.Println(cnama)
	fmt.Println(len(cnama))
	fmt.Println(cap(cnama))
	/* -append di gunakan untuk menambahkan slice baru dengan menambahkan data di posisi terakhir
	   -jika kapasitas sudah penuh maka akan menambahkan array baru
	   -jika len < dari cap ketika di append, maka akan terjadi replace
	*/
	fmt.Println("penggunaan append (kapasitas penuh)")
	abuff := append(bnama, "toge") // menambahkan data di abuff
	fmt.Println(abuff)
	abuff[0] = "rubah" // merubah data index ke 0 di abuff
	fmt.Println(abuff)
	fmt.Println(bnama) //data tidak berubah
	fmt.Println(nama)  //data tidak berubah
	/* ketika di appand , data di slice abuff akan bertambah,dan ketika di rubah abuff akan berubah
	tapi ketika sumbernya di print (bnama dan nama) tidak terjadi perubahan apa pun. ini terjadi karena
	len = cap (kapasitasnya sudah penuh) jadi golang membuat array baru*/
	fmt.Println("penggunaan append (kapasitas tidak penuh)")
	bbuff := append(cnama, "tomat")
	fmt.Println(bbuff)
	bbuff[0] = "rubah"
	fmt.Println(bbuff)
	fmt.Println(cnama)
	fmt.Println(nama)
	/*
		membuat slice baru , ini lebih aman di karenakan arraynya di sembunyikan di belakang slice\
		make([]tipe data,jumlah elemen,capacity)
	*/

	x := make([]string, 2, 5)
	x[0] = "naga"
	x[1] = "langit"
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	/* copy slice usahakan ukuran len dan cap nya sama */
	y := make([]string, 2, 5)
	copy(y, x)
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

}
