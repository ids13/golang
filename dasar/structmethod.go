package main

import "fmt"

type customer struct {
	nama, kerjaan string
	umur          int
}

func (a customer) say() {
	fmt.Println("hallo ", a.nama)
}

func main() {
	var data customer
	data.nama = "naga"
	data.kerjaan = "karyawan"
	data.umur = 30

	data.say()

	data2 := customer{
		nama:    "joko",
		kerjaan: "pedagang",
		umur:    40,
	}
	data2.say()
}
