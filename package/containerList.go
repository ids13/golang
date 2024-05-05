package main

import (
	"container/list"
	"fmt"
)

// link list ini di bagian akhir nya nil jadi jangan terlalu sering next jika tidak tau ada berapa datanya
func main() {
	data := list.New()    //membuat list
	data.PushBack("naga") //memasukan data ke belakang
	data.PushBack("langit")
	data.PushBack("awan")
	for e := data.Front(); e != nil; e = e.Next() { //cek data di bagian depan, jika tidak kosong makan next
		fmt.Println(e.Value)
	}
}
