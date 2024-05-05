package main

import "fmt"

func main() {
	/* for loop adalah perulangan yang di gunakan jika jumlah perulangan di ketahui dan
	juga untuk melakukan perulangan pada array dan slice */
	angka := 0

	// hanya kondisi
	for angka < 5 {
		fmt.Println("angka", angka)
		angka++
	}
	// meggunakan argumen
	for i := 0; i < 5; i++ {
		fmt.Println("hai", i)
	}

	buah := [...]string{"apel", "mangga", "jambu"}
	/* penggunaan range di for untuk melakukan perulangan pada map,slice,array,string
	   contoh di bawah."index menujukan key nya (nama variable nya bisa bebas,tidak harus index).
	   dan val adalah variable penampung untuk nilainya
	   untuk pemisah antara index dan val menggunakan tanda KOMA bukan titik koma"
	   jika salah satu di antara index / var tidak di perlukan maka harus di ganti dengan tanda garis bawah "_" */
	for index, val := range buah {
		fmt.Println("index = ", index, "value = ", val)
	}

	numbers := []int{1, 2, 3, 4, 5}
	// Menggunakan range tanpa menggunakan indeks atau nilai
	for range numbers {
		fmt.Println("Iterasi")
	}
	// contoh penggunaan break dan continue
	for i := 0; i < 10; i++ {
		if i == 4 {
			continue
		} else if i == 7 {
			break
		} else {
			fmt.Println("hai ", i)
		}
	}
	// pengulangan range dana angka yg sudah di tentukan
	for range 5{
		fmt.Println("hallo")
	}

	/* 
	ini untuk infinite looping
	for{

	}
	*/
}
