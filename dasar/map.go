package main

import "fmt"

func main() {
	/* array dan slice menggunakan index untuk mengakses data
	   map bisa kita tentukan tipe dan vlaue
	   sederhananya map adalah kunci dan value
	   kata kunci nya tidak boleh sama, jika ada yang sama makan akan tertimpa oleh data yang lebih baru
	*/
	x := map[string]string{
		"negara":  "indonesia",
		"kota":    "jakarta",
		"makanan": "kerak telor",
	}
	fmt.Println(x)
	/* ketika di tampilkan.urutan bisa saja acak. tapi itu tidak penting
	penggunaan map untuk mengakses nilai menggunakan kunci */
	fmt.Println(x["negara"])
	// penambahan data, kunci harus unik
	x["monumen"] = "monas"
	fmt.Println(x)
	delete(x, "monumen")
	fmt.Println(x)
	// deklarasi map
	y := make(map[string]string)
	y["makanan"] = "dodol"
	fmt.Println(y)
}
