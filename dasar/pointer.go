package main

import "fmt"

type alamat struct {
	kota string
	no   int
}

func main() {
	alamatA := alamat{"sumedang", 3}
	alamatB := &alamatA      //tanda & menunjukan pointer
	alamatB.kota = "bandung" //field struct tidak perlu tanda bintang,
	fmt.Println(alamatA)
	fmt.Println(alamatB)
	*alamatB = alamat{"baru", 10}
	fmt.Println(alamatA)
	fmt.Println(alamatB)

	var nilai *int = new(int) // new(tipedata)
	*nilai = 20
	fmt.Println(*nilai)

	val1 := 80
	val2 := &val1
	fmt.Println(*val2)
}
