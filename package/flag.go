package main

import (
	"flag"
	"fmt"
)

/* flag di gunakan untuk parsing command line argumen */
func main() {
	/* untuk menjalankan gunakan
	go run flag.go
	jika -host dan -usernamenya tidak di set ketika mengeksekusi
	maka akan di set nilai defaultnya
	dan jika salah memasukan option (menggunakan option selain -host dan -username)
	maka pesan error/help(bisa juga menggunakan -h / -help) akan di tampilkan */
	host := flag.String("host", "localhost", "masukan nama host anda")
	username := flag.String("user", "user", "masukan nama user anda")
	flag.Parse()
	/* hasil dari flag bentuk nya alamat memmory
	jadi untk mengakses nilainya harus menggunakan tanda bintang */
	fmt.Println(*host, *username)
}
