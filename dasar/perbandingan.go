package main

import "fmt"

func main() {
	s1 := "dragon"
	s2 := "fly"
	buff := s1 == s2
	fmt.Println("kata ", buff)
	i1 := 2
	i2 := 3
	fmt.Println(i1 < i2)
	fmt.Println(i1 > i2)
	fmt.Println(i1 <= i2)
	fmt.Println(i1 >= i2)
	fmt.Println(i1 == i2)
	fmt.Println(i1 != i2)
}
