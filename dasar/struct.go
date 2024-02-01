package main

import "fmt"

/* struct adalah templete data,menggabungkan nol atau lebih tipe data(karena jika array,tipe datanya harus sama)
simpelnnya struct adalah kumpulan dari field*/
type customer struct {
	nama, kerjaan string
	umur          int
}

func main() {
	var data customer
	data.nama = "naga"
	data.kerjaan = "karyawan"
	data.umur = 30

	fmt.Println(data)

	data2 := customer{
		nama:    "joko",
		kerjaan: "pedagang",
		umur:    40,
	}
	fmt.Println(data2.nama)
}
