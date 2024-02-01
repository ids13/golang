package main

import "fmt"

func main() {
	/* jika hasil expresinya menghasilkan true,maka block programnya akan di eksekusi */
	a := "lina"

	if a == "hari" {
		fmt.Println("hallo hari")
	} else if a == "lina" {
		fmt.Println("hallo lina")
	} else {
		fmt.Println("maaf siap ya")

	}
	// shortener if
	if nilai := 8; nilai < 7 {
		fmt.Println("nilai di bawah kkm")
	} else {
		fmt.Println("anda lulus")
	}
}
