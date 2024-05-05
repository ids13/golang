package main

import (
	"fmt"
	"reflect"
)

// Struct contoh
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Mengambil Informasi Tipe Data dan Nilai
	var x float64 = 3.14
	t := reflect.TypeOf(x)  // Mengambil tipe dari nilai x
	v := reflect.ValueOf(x) // Mengambil nilai dari variabel x
	fmt.Println("1. Tipe x:", t)
	fmt.Println("2. Nilai x:", v)

	// Memeriksa Tipe dan Nilai
	kind := t.Kind()               // Mengambil jenis tipe nilai x
	val := v.Interface().(float64) // Mengambil nilai x dalam bentuk interface{}
	fmt.Println("3. Jenis Tipe x:", kind)
	fmt.Println("4. Nilai x dalam interface{}:", val)

	// Mendapatkan Informasi Struktur dan Tag
	tPerson := reflect.TypeOf(Person{})
	numField := tPerson.NumField() // Mengambil jumlah field pada struktur Person
	fmt.Println("5. Jumlah Field Person:", numField)
	for i := 0; i < numField; i++ {
		field := tPerson.Field(i) // Mengambil informasi tentang field ke-i
		tag := field.Tag          // Mengambil tag yang terkait dengan field
		fmt.Printf("   - Field ke-%d: %s, Tag: %s\n", i+1, field.Name, tag.Get("json"))
	}

	// Mendapatkan Informasi Tambahan
	numMethod := tPerson.NumMethod() // Mengambil jumlah metode yang dimiliki oleh tipe Person
	fmt.Println("6. Jumlah Metode Person:", numMethod)

	// Manipulasi Nilai
	var y string = "hello"
	vY := reflect.ValueOf(&y).Elem() // Mendapatkan nilai variabel y yang dapat diubah
	if vY.CanSet() {
		vY.SetString("world") // Mengubah nilai variabel y menjadi "world"
	}
	fmt.Println("7. Nilai y setelah diubah:", y)

	// Konversi Tipe
	var a float64 = 3.14
	vA := reflect.ValueOf(a)                   // Mengambil nilai dari variabel a
	aInt := vA.Convert(reflect.TypeOf(int(0))) // Mengkonversi nilai variabel a menjadi tipe int
	fmt.Println("8. aInt:", aInt.Interface())
}
