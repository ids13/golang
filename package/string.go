package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "naga langit petir"
	//apakah di string ada kata langit
	fmt.Println(strings.Contains(text, "langit"))
	//memisahkan dengan sparatornya spasi,nilai returnya adalah slice
	fmt.Println(strings.Split(text, " "))
	//semua jadi huruf kecil
	fmt.Println(strings.ToLower(text))
	//semua jadi huruf besar
	fmt.Println(strings.ToUpper(text))
	//huruf besar di setiap awal kata
	fmt.Println(strings.Title(text))
	// menghapus karakter tertentu di awal dan akhir string,contoh di bawah akan menghapus spasi
	fmt.Println(strings.Trim("     naga     ", " "))
	text2 := "halo halo halo dunia"
	//mengganti string dengan jumlah n kali
	fmt.Println(strings.Replace(text2, "halo", "hai", 1))
	//mengganti semua string
	fmt.Println(strings.ReplaceAll(text2, "halo", "hai"))
	//mengulang char
	fmt.Println(strings.Repeat("A", 10))
}
