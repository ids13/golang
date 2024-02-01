package main

import "fmt"

func main() {
	a := 6
	b := 7
	c := 2
	var buff int
	//penjumlahan
	buff = a + b
	fmt.Println(buff)
	//pengurangan
	buff = a - b
	fmt.Println(buff)
	//pengkalian
	buff = a * b
	fmt.Println(buff)
	//pembagian
	buff = a / c
	fmt.Println(buff)
	//sisa pembagian
	buff = b % c
	fmt.Println(buff)

	//asign argument
	buff = 100
	buff += 10
	fmt.Println(buff)
	buff -= 20
	fmt.Println(buff)

	//unary oprator (++,--,!)
	buff++
	fmt.Println(buff)
	status := true
	fmt.Println(!status)
}
