package main

import "fmt"

/* struct adalah templete data,menggabungkan nol atau lebih tipe data(karena jika array,tipe datanya harus sama)
simpelnnya struct adalah kumpulan dari field*/
type customer struct {
	nama, kerjaan string
	umur          int
}

type orang struct {
	nama string
	umur int
}

type murid struct {
	orang
	kelas int
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

	data3 := customer{"andi", "kuliah", 20}
	fmt.Println(data3)
	//embed struct
	var data4 murid
	data4.kelas = 12
	data4.nama
	data4.umur = 17

	fmt.Println(data4.kelas)
	fmt.Println(data4.nama)
	fmt.Println(data4.umur)

	/*
		Go menyediakan akses langsung ke bidang-bidang yang ditanamkan
		tanpa perlu menyebutkan nama struct yang ditanamkan.
		Ini karena Go memperkenalkan mekanisme yang disebut "promotion" (peningkatan) untuk bidang-bidang yang ditanamkan.
		Promotion berarti bahwa bidang-bidang dari struct yang ditanamkan secara otomatis "naik" ke struct yang menanamkannya.
		Ketika Anda mengakses bidang-bidang yang ditanamkan,
		Go akan mencari bidang tersebut terlebih dahulu di struct yang menanamkannya.
		Jika bidang tersebut tidak ditemukan di struct yang menanamkannya,
		Go akan mencari di struct yang ditanamkan.
		Dalam contoh Anda, ketika Anda menulis data4.umur, Go akan mencari bidang umur di dalam struct murid.
		Karena bidang umur tidak ditemukan di dalam struct murid,
		Go akan melihat ke dalam struct yang ditanamkan, yaitu orang, dan menemukan bidang umur di sana.
		Oleh karena itu, Anda bisa langsung menulis data4.umur
		untuk mengakses bidang umur tanpa perlu menuliskan data4.orang.umur(kecuali ada nama field yang sama).
		jika memiliki nama field yang sama (contohnya struct murid juga memiliki field umur)
		maka cara aksesnya data4.umur dan data4.orang.umur , ini supaya tidak ambigu
	*/
}
