package main

import "fmt"

func main() {
	squareOf2 := func() int {
		return 2 * 2
	}() // ini artinya iifi , fungsinya akan langsung di jalan kan dan megembalikan return int.
	fmt.Println(squareOf2) // jka tidak menggunakan iifi squareOf2 akan bertipe fungsi, dan jika menggunakna iifi squareof2 akan bertipe int

	// jika tidak menggunakna iifi
	squareOf := func() int {
		return 2 * 2
	}
	fmt.Println(squareOf())
}
