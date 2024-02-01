package main

import (
	"fmt"
)

/* di golang hanya beberapa tipe data saja yang bisa di set nil
interface,slice,pointer,map,channel,function
di function kita tidak bisa mengeset nilai nil seperti pada lainnya

	var myFunc func()

    // Memanggil fungsi dengan variabel fungsi yang belum diinisialisasi
    panggilFungsi(myFunc) // Output: Fungsi belum diinisialisasi (pengecekan)

    // Menetapkan fungsi ke variabel
    myFunc = func() {
        fmt.Println("Halo dari myFunc!")
    }

    // Memanggil fungsi dengan variabel fungsi yang sudah diinisialisasi
    panggilFungsi(myFunc) // Output: Halo dari myFunc!

*/

func main() {
	var ptr *int              //deklarasi pointer
	var nilSlice []int        //deklarasi slice
	var nilMap map[string]int //deklarasi map
	var nilChannel chan int   //deklarasi chanel
	var nilFunc func()        //deklarasi fungsi
	//pengecekan
	if ptr == nil {
		fmt.Println("Pointer is nil")
	} else {
		fmt.Println("Pointer is not nil")
	}
	if nilSlice == nil {
		fmt.Println("slice is nil")
	} else {
		fmt.Println("slice is not nil")
	}
	if nilMap == nil {
		fmt.Println("map is nil")
	} else {
		fmt.Println("map is not nil")
	}
	if nilChannel == nil {
		fmt.Println("channel is nil")
	} else {
		fmt.Println("channel is not nil")
	}
	if nilFunc == nil {
		fmt.Println("func is nil")
	} else {
		fmt.Println("func is not nil")
	}
	//pengisian nilai
	val := 42
	ptr = &val
	nilSlice = []int{1, 2, 3}
	nilMap = map[string]int{"a": 1, "b": 2}
	nilChannel = make(chan int)
	nilFunc = func() {
		fmt.Println("Hello from not-nil function")
	}

	if ptr == nil {
		fmt.Println("Pointer is nil")
	} else {
		fmt.Println("Pointer is not nil")
	}
	if nilSlice == nil {
		fmt.Println("slice is nil")
	} else {
		fmt.Println("slice is not nil")
	}
	if nilMap == nil {
		fmt.Println("map is nil")
	} else {
		fmt.Println("map is not nil")
	}
	if nilChannel == nil {
		fmt.Println("channel is nil")
	} else {
		fmt.Println("channel is not nil")
	}
	if nilFunc == nil {
		fmt.Println("func is nil")
	} else {
		fmt.Println("func is not nil")
	}

}
