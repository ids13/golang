package main

import "fmt"

func main() {
	/*
	   && dan
	   || atau
	   ! negasi/kebalikan
	*/

	nilai1 := 60
	nilai2 := 80
	buff1 := nilai1 >= 70
	buff2 := nilai2 >= 70
	fmt.Println(buff1 && buff2)

}
