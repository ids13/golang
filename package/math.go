/* package math berisi constanta dan fungsi matematika */
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.4))  // membulatkan ke atas/ke bawah tergantung paling dekat ke mana
	fmt.Println(math.Floor(1.4))  // membulatkan ke bawah
	fmt.Println(math.Ceil(1.4))   // membulatkan ke atas
	fmt.Println(math.Max(12, 20)) // mengambil yang paling besar
	fmt.Println(math.Min(12, 20)) // mengambil yangn paling kecil
}
