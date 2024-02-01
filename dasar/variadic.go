package main

import "fmt"

/* parameter pada fngsi pasa fungsi di posisi paling terakhir memiliki kemampuan sebagai vargs
yaitu bisa menerima data lebih dari satu input(gampangnya mirip seperti array dengan penggunaan titik tiga)
ini hanya bisa di buat satu dan harus di letakan di paling terakhir*/

/* ini contoh penggunaan variadic function,variadic harus di tempatkan di paling kanan.
jika ingin menambahkan parameter lain, harus di tambahkan di sebelumnya
contoh : sumall(par string,num...int) */
func sumall(num ...int) int {
	var total int
	for _, a := range num {
		total += a
	}
	return total
}

func main() {

	//variadic menggunakan array
	tot := sumall(11, 22, 33, 44, 55)
	fmt.Println(tot)

	//varidic menggunakan slice
	angka := []int{1, 2, 3, 4, 5}
	t := sumall(angka...) //harus menambahkan tiga titik,untuk memberitahu variadic parameter
	fmt.Println(t)
}
