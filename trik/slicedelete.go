package main

import "fmt"

func main() {
	// Membuat slice contoh
	mySlice := []int{1, 2, 3, 4, 5}

	// Menghapus elemen dengan index 2 dari slice menggunakan copy
	indexToDelete := 2
	fmt.Println(mySlice[2:])
	fmt.Println(mySlice[3:])
	copy(mySlice[indexToDelete:], mySlice[indexToDelete+1:])
	fmt.Println(mySlice)
	mySlice = mySlice[:len(mySlice)-1]

	// Menampilkan slice setelah penghapusan
	fmt.Println(mySlice)
}
