package main

import "fmt"

func halo(nama string) string {
	return "hallo " + nama //tanda plus yang di tempatkan di dekat string fungsinya untuk penggabungan string
}

func main() {
	h := halo("naga")
	fmt.Println(h)
}
