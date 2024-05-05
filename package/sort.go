package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	umur int
}

type SortBy []person

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].umur < a[j].umur }

func main() {
	people := []person{
		{"aan", 1},
		{"baba", 5},
		{"caca", 2},
	}
	/* people ini kan tipe data slice bertipe struct person,sedangkan SortBy adalah
	alias jadi harus di konversi dulu dengan SortBy(people) */

	sort.Sort(SortBy(people))
	fmt.Println(people)

	kata := []string{"apel", "jeruk", "pisang", "melon"}
	angka := []int{0, 2, 4, 5, 3, 2, 1, 4, 3}
	koma := []float64{1.2, 3.2, 1.4, 3.1, 3.2, 3.3}
	kata2 := []string{"apel", "jeruk", "pisang", "melon"}
	// perbedaan .strings dan .steringslice adalah stringslice di tampung dulu

	penampung := sort.StringSlice(kata2)
	sort.Strings(kata)
	sort.Ints(angka)
	sort.Float64s(koma)

	fmt.Println(penampung)
	fmt.Println(kata)
	fmt.Println(angka)
	fmt.Println(koma)

}
