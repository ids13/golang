package main

import "fmt"

func main() {
	/*beberapa cara untuk menginisialisasi variable
	  di dalam golang variable yang di deklarasikan harus di gunakan,jika tidak maka akan menimbulkan error
	*/
	var name string
	name = "dragon"
	fmt.Println(name)
	//jika data langsung di isikan ke dalam variable maka tipe datanya bisa tidak di sebutkan karen golang akan membacanya secara otomatis
	var name2 = 123
	fmt.Println(name2)
	// := di gunakan untuk deklarasi variable dan langsing mengisi. lakukan ini sekali karen deklarsi tidak boleh berkali kali
	name3 := "fly"
	fmt.Println(name3)
	//multiple deklarasi variable
	var (
		data1 string
		data2 = 10
	)
	data1 = "multi"
	fmt.Println(data1)
	fmt.Println(data2)
}
