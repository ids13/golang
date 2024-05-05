Inisialisasi project
go mod init nama-modulnya
    -nama modul jangan ada spasinya
    -untuk mempermudah nama modul sesuaikan dengan nama foldernya
    -gunakan nama-modul ini di bagian paling atas code (contoh : package main(ini defaultnya di golang))

Menjalankan file golan
go run namafile.go

compile program
go build

varable
| nama   | keterangan                                        |
| ------ | ------------------------------------------------- |
| int    | sama dengan int32 atau int64 (tergantung nilai)   |
| int8   | -128 ↔ 127                                        |
| int16  | -32768 ↔ 32767                                    |
| int32  | -2147483648 ↔ 2147483647                          |
| int64  | -9223372036854775808 ↔ 9223372036854775807        |
| uint   | sama dengan uint32 atau uint64 (tergantung nilai) |
| uint8  | 0 ↔ 255                                           |
| uint16 | 0 ↔ 65535                                         |
| uint32 | 0 ↔ 4294967295                                    |
| uint64 | 0 ↔ 18446744073709551615                          |
| byte   | sama dengan uint8                                 |
| rune   | sama dengan int32                                 |

inisialisasi
- var nama string
- var nama = "nilai"
- nama := "nilai"
- nama1,nama2 := "nilai1","nilai2"
-  var (
   nama1="nilai1"
   nama2="nilai2"
   )

constanta
constanta boleh tidak di gunakan namun tidak boleh di rubah isinya
- const nama string = "nilai"
- const nama := "nilai
- const (
   nama1="nilai1"
   nama2="nilai2"
   )

konversi tipe data
jika melebihi kapasitas, maka akan balik lagi ke belakang.
- string(tipeByte)
- int(tipeIntBerapa)

type declaration
kemampuan membuat tipe data baru dari tipe data yang sudah ada
type noKTP string

oprasi matematika
| oprator | keterangan |
| ------- | ---------- |
| +       | tambah     |
| -       | kurang     |
| *       | kali       |
| /       | bagi       |
| %       | sisa bagi  |
ini untuk tipe data numer

argumen assigment
| oprator | keterangan |
| ------- | ---------- |
| a=a+x   | a+=x       |
| a=a-x   | a-=x       |
| a=a*x   | a*=x       |
| a=a/x   | a/=x       |
| a=a%x   | a%=x       |

unary oprator
| oprator | keterangan         |
| ------- | ------------------ |
| ++      | naikan satu angka  |
| --      | kurangi satu angka |
| +       | positif            |
| -       | negative           |
| !       | boolean kebalikan  |

oprator perbandingan
| oprator | keterangan              |
| ------- | ----------------------- |
| >       | lebih besar             |
| <       | lebih kecil             |
| >=      | lebih besar sama dengan |
| <=      | lebih kecil sama dengan |
| ==      | sama dengan             |
| !=      | tidak sama dengan       |

oprasi boolean
| oprator | keterangan                   |
| ------- | -----------------------------|
| &&      | benar jika keduanya true     |
| ||      | benar jika salah satunya tru |
| !       | kebalikan dari nilai         |

oprasi bitwise
| Operator | Keterangan  | Contoh   | Hasil (Desimal) |
| -------- | ----------- | -------- | --------------- |
| &        | AND bitwise | `5 & 3`  | 1               |
| \|       | OR bitwise  | `5 \| 3` | 7               |
| ^        | XOR bitwise | `5 ^ 3`  | 6               |
| <<       | Left Shift  | `5 << 1` | 10              |
| >>       | Right Shift | `5 >> 1` | 2               |


array
sekumpulan data yang bertipe sama,elemen pertama array di wakili index 0
inisialisasi
- var array [3]string 
- arr := [...]string{"nilai1", "nilai2","nilai3","nilai4"}
  penggunaan
  - array[0] = "nilai1"
  - array = {"nilai1","nilai2","nilai3"}
  - array = [2][3]int{{1, 2, 3}, {1, 2, 3}}
  
| fungsi             | keterangan                     |
| ------------------ | ------------------------------ |
| len(arr)           | panjang array                  |
| arr[index]         | mendapatkan data posisi index  |
| arr[index] = nilai | merubah data data posisi index |

slice
potongan dari tipe data array,ukuran bisa berubah,array dan slice saling terkoneksi(satu rubah semua rubah)
a := [3]int{1,2,3} ini array
a := [...]int{1,2,3} ini array
a := []int{1,2,3} ini slice

inisialisasi
- var emptySlice []int  (ini deklarasi)
- initializedSlice := []int{1, 2, 3}
- makeSlice := make([]int, 5,5)
  pengunaan
  -  arr := [...]string{"nilai1", "nilai2","nilai3","nilai4"}
     s := arr[1:3]
     ouput nya [nilai 2 nilai 3]

| membuat         | keterangan                                                   |
| --------------- | ------------------------------------------------------------ |
| array[low:high] | membuat slice dari array index low sampai index sebelum high |
| array[low:]     | membuat slice dari array index low sampai akhir array        |
| array[:high]    | membuat slice dari awal array sampai index sebelum high      |
| array[:]        | membuat slice dari array index 0 sampai akhir array          |


| fungsi                            | keterangan                                                                                                    |
| --------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| len(slice)                        | panjang slice(bukan panjang arraynya)                                                                         |
| cap(slice)                        | kapasitas slice(si hitung dari pointer penunjuk data pertamanya sampai akhir array)                           |
| append(slice,data)                | membuat slice baru dengan menambah data ke posisi terakhir,jika kapasitas penuh .maka akan membuat array baru |
| make([typeData],panjang,kapasitas | membuat slice baru                                                                                            |
| copy(tujuan,sumber)               | menyalin slice dari sumber ke tujuan                                                                          |

map
sekumpulan tipe data sama, yang jenis tipe data index nya bisa di tentukan
sederhananya map adalah kumpulan key-value
beda dengan slice dan array,map bisa dimasukan data sebanyak banyak nya

inisialisasi
- var emptyMap map[string]int (ini deklarasi)
- initializedMap := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
- makeMap := make(map[string]bool)
  
    penggunaan
    - data := make(map[string]string)
      data["nama"] = "naga"
      data["alamat"] = "subang"

| oprasi                              | keterangan                                            |
| ----------------------------------- | ----------------------------------------------------- |
| len(map)                            | jumlah data di map                                    |
| map[key]                            | mengambil data di map dengan key                      |
| map[key] = value                    | mengubah data di key,jika belum ada maka akan di buat |
| make(map[tipeDataKey]tipeDataValue) | membuat map baru                                      |
| delete(map,key)                     | menghapus data di map                                 |

if
jika kondisinya bernilai true makan if akan di eksekusi ,dan jika false maka else yang akan di eksekusi
else if menambahkan blok if tambahan

if x == y {
   
}else if x == z{

}else{

}

short statment if

if a := 10;a < 100 {

}

switch
melakukan pengecekan ke kondisi dalam suatu variable. jika nilainya true maka case nya akan di eksekusi dan jika tidak ada di case maka default akan di eksekusi

sederhana
switch variable{
   case "a":
      //lakukan sesuatu
   case "b":
      //lakukan sesuatu
   default:
      //lakukan sesuatu
}
//ini sperti berpindah(switch) ke variable kemudian membandingkan nilainya dengan yang ada di case

sort statment
switch a := 3; a < 10{
   case true:
      //lakukan sesuatu
   case false:
}

tanpa switch expression
switch {
   case variable < 10:
      //lakukan sesuatu
   case variable < 100:
      //lakukan sesuatu
   default:
      //lakukan sesuatu
}
//di sini variable nya tidak di switch oleh sebab itu, program akan membandingkan nilai expression dari setiap casenya

for 
melakukan perulangan

sederhana
for a < 10{
   //lakukan sesuatu
   a++
}

dengan statment
for a := 1 ; a < 10;a++{
   //lakukan sesuatu
}

infinite looping
for {
    // Lakukan sesuatu yang relevan di sini
}

menggunakan range (mulai di versi 1.22)
func main() {
	for i := range 5 {
		fmt.Println(i)
	}
}


range(iterasi)
numbers := [5]int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
//melakukan iterasi terhadap data collection seperti array,slice,map

break dan continue 
break : menghentikan perulangan
continue : menhentikan perulangan yang berjalan,kemudian lanjut ke perulangan berikutnya

contoh
for a := 1 ; a < 10;a++{
   if a == 3 {
      break
   }
}

fungsi
fungsi adalah blok kode yang memiliki nama dan dapat dipanggil untuk melakukan tugas tertentu.tidak boleh ada tanda - / _ pada nama fungsi (gunakan kapital di awal kata jika perlu pemisah)

func namaFungsi (){

}
fung main(){
   namaFungsi()
}

parameter
nilai yang diterima oleh fungsi saat dipanggil


func namaFungsi (a int,b string){

}
fung main(){
   namaFungsi(dataTypeInt,dataTypeString)
}

return
pernyataan dalam Golang yang digunakan untuk mengembalikan nilai dari suatu fungsi

func namaFungsi () tipeDataReturn{
   return dataTypeReturn
}
fung main(){
   x := namaFungsi()
}

multi return
func namaFungsi () (tipeDataReturn1,tipeDataReturn2){
   return dataTypeReturn1,dataTypeReturn2
}
fung main(){
   x,y := namaFungsi()
   //jika returnnya ingin di abaikan gunakan tanda garis bawah (_)
}

named return
func namaFungsi () (var1,var2 typeDataReturn){
   var1 = "halo"
   var2 = "dunia"
   return
}
fung main(){
   x,y := namaFungsi()
   //jika returnnya ingin di abaikan gunakan tanda garis bawah (_)
}

variadic
parameter paling terakhir bisa di jadikan varargs
artinya datanya bisa menerima lebih dari satu input,anggap saja sperti array. posisi parameter nya wajib di paling ujung. gunakan tanda titi tiga kali sebelum tipe data parameternya

contoh
func namaFungsi (num ...int){
   for _,i := range num {
      fmt.Println(i)
   }
}
func main(){
   namaFungsi(1,2,3,4,5,6,7,8,9)
   //jika datanya sudah berupa slice,gunakan titik tiga kali
   namaFungsi(namaSlice...)
}

as value
func namaFungsi(){

}
func main(){
   x := namaFungsi
   x()
}

as parameter
func namaFungsi(var1 string,var2 func(string)string){

}
func contoh (string)string{

}
func main(){
x := contoh
namaFungsi("string",x(string))
}

anonymous
Anonymous function di Golang adalah sebuah fungsi tanpa nama yang dapat didefinisikan dan langsung digunakan di dalam suatu program tanpa perlu memberikan nama ke fungsi tersebut. Anonymous function umumnya digunakan dalam situasi di mana kita hanya memerlukan fungsi sederhana untuk digunakan secara lokal

func namaFungsi (var1 string,var2 func(string)bool){

}
func main(){
   x := func(string)bool{return true}
   namaFungsi("string",x)
   //atau
   namaFungsi("string",func(string)bool{
      return tru})
}

recursive
Fungsi rekursif (recursive function) adalah fungsi yang memanggil dirinya sendiri selama proses eksekusi.

func factorial(n int) int {
    // Kasus dasar: faktorial dari 0 atau 1 adalah 1
    if n <= 1 {
        return 1
    }
    // Panggilan rekursif untuk menghitung faktorial
    return n * factorial(n-1) //contoh recursive
}

closures
kemampuan fungsi berinteraksi dengan data data disekitarnya dalam lingkup yang scope yang sama
contoh: jika ada fungsi di dalam fungsi main . maka fungsi tersebut bisa mengakses data (contohnya variable) di luar fungsi tersebut selama data itu masih berada di dalam fungsi main

defer
Defer adalah kata kunci dalam Go yang digunakan untuk menunda (defer) eksekusi suatu pernyataan atau fungsi hingga fungsi yang berisi defer tersebut selesai dieksekusi, atau hingga blok program tempat defer tersebut ditempatkan selesai dieksekus.defer akan di eksekusi meskipun programnya error.

defer namaFungsi()
//tinggal tempatkan ini di fungsi lain/di fungsi main

panic
panic adalah fungsi yang di gunakan untuk menghentikan program

panic("aplikasi error")

recover
recover adalah fungsi yang di gunakan untuk menangkap data panic.

message := recover()

komentar
// untuk single line
/**/ untuk multi line

struct
struct templete data yang di gunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan

pembuatan
type namaStruct struct {
	var1 int
	var2,var3 string
}

inisialisasi
var cnth namaStruct
cnth := namaStruct{1, "naga", "langit"}
cnth := namaStruct{
		var1: 1,
		var2: "text1",
		var3: "text2",
	}
akses data
cnth.var1 = 1 
cnth.var2 = "text1"
cnth.var3 = "text2"

kita juga bisa membuat method di dalam struct

type namaStruct struct {
    // Field-field di sini
}
// Definisikan sebuah method terkait dengan kendaraan
// menggunakan pointer agar ketika terjadi perubahan data di struct namaStruct di dalam fungsi metodeOprasi,namaStruct juga ikut terpengaruh
func (k *namaStruct) metodeOperasi() {
    // Implementasi method di sini
}
func main() {
    // Buat sebuah instance dari kendaraan
    namaStructSaya := namaStruct{}
    // Panggil method pada instance
    namaStructSaya.metodeOperasi()
}

interface
adalah tipe data yang abstrak,biasanya di gunakan sebagai kontrak

// kita membuat kontrak fungsi gettype di interface, jadi siapapun struct yang memiliki method dengan format gettype maka secara otomatis terikant kontrak dengan interface ini
type mahluk interface {
	getType() string
}
type manusia struct {
	nama string
}
// ini contoh penjalinan kontrak
func (m manusia) getType() string {
	return m.nama
}
//contoh implementasi penggunaan data di interface
func say(x mahluk) {
	fmt.Println(x.getType())
}
func main() {
	cnth := manusia{"andi"}
	say(cnth)
}
//jadi dengan interface , kita bisa membuat struct mahluk lainnya (hewan /tumbuha),kemudian membuat method dengan format gettype di struct tersebut supaya terjalin kontrak dengan interface mahluk.dengan begini kita bisa memanggil fungsi say dan memasukan struct yang sudah terikat dengan kontrak interface mahluk sebagai parameternya.

interface kosong
interface yang tidak memiliki deklarasi method satupun,hal ini membuat secara otomatsi semua tipe data akan menjadi implementasinya. jadi bisa di gunakan sebagai tipe data,return,parameter method/fungsi
any adalah alias untuk interface kosong.

contoh sebagai parameter
func say(data interface{}) {
	fmt.Println(data)
}
func main() {
	say(1234)
	say("test")
}

nil
saat membuat variable dengan tipe data tertentu,maka akan di buat default valuenya. nil ini sama dengan data kosong,hanya bisa di gunakan untuk beberapa tipe data seperti interface,function,map,slice,pointer dan channel

contoh
func say(nama string) interface{} {
	if nama == "anjing" {
		return nil
	} else {
		return "hallo " + nama
	}
}
func main() {
	fmt.Println(say("anjing"))
}

type assertion
kemampuan merubah data menjadi tipe data yang kita ingin kan,sering di gunakan ketika bertemu dengan interface kosong.

	var a interface{}
	a = "hallo"
	// type assertion
	b := a.(string)
	//pengecekan tipe string
	// jika berhasil type assertion berhasil str akan berisi string,err akan berisi true
	// jika berhasil type assertion berhasil str akan berisi 0,err akan berisi false
	str, err := a.(int)
	//type assertion (type) hanya berlaku di switch
	switch a.(type) {
	case int:
		fmt.Println("ini int")
	case string:
		fmt.Println("ini string")
	}

pointer
Pointer adalah variabel yang berisi alamat memori dari nilai variabel lain.
deklarasi
var ptr *int
ptr := new(int) //perintah new ini hanya di gunakan di dat pointer
Operator ampersand (&) untuk mendapatkan alamat memmory suatu variable
Operator asterisk (*) digunakan untuk mengakses nilai yang disimpan di alamat yang ditunjuk oleh pointer.
Penggunaan * sebelum variabel pointer:
melakukan dereferensi, yang berarti mengakses nilai yang disimpan di alamat memori yang ditunjuk oleh pointer
Penggunaan * dalam deklarasi pointer:
untuk menunjukkan bahwa variabel tersebut adalah sebuah pointer.
Penggunaan tanpa * saat mengakses field struct melalui pointer:
dari struct yang ditunjuk oleh pointer tersebut. Tidak perlu menggunakan tanda bintang (*) karena Go menangani dereferensi secara otomatis.(berlaku untuk satu tingkat * ,jika ** kita harus menggunakan satu *)

parameter fungsi
func say(nama *string){

}
func main(){
   say(&var)
}

sangat di rekomendasikan menggunakan pointer di method, jadi tidak boros memmory
type Man struct {
	name string
}
func (man *Man) nikah() { //penggunaan pointer untuk terikat ke struct,dan untuk aksed datanya tidak perlu tanda * karena mengakses field struct.
	man.name = "mr." + man.name
}
func main() {
	v := Man{"naga"} //menyimpan pointer struct man, dan megeset isi var nama di struct man
	v.nikah()
	fmt.Println(v.name)
}

package
di golang kita membuat paket untuk merapihkan kode file
contoh 
dasar(folder)
   go.mod
   utama.go
   helper(folder)
      helper.go
-hanya boleh ada satu fungsi main di satu project golang
-nama packade samakan dengan nama foldernya
-untuk melihat nama project lihat di go.mod
-jadi ketika di import di file nya seperti contoh di atas gunakan import "dasar/helper" di file utama.go
untuk akses gunakan helper.Namafungsinya()

access modifier
mekanisme dalam pemrograman yang mengatur aksesibilitas dari anggota suatu kelas, struktur, atau objek dalam sebuah program.
Identifier yang diawali dengan huruf besar adalah publik dan dapat diakses dari luar paket.
Identifier yang diawali dengan huruf kecil adalah pribadi dan hanya dapat diakses dari dalam paket.

initialization
-kita bisa membuat fungsi yang akan dijalankan ketika paket di akses
-contoh kasus, ketika membuat data base kita bisa bisa membuat fungsi yang otomatis menginisialisasi koneksi ke data base .
-mengunakan init

contoh
// Fungsi init akan dieksekusi secara otomatis sebelum fungsi main dijalankan
func init() {
    fmt.Println("Ini adalah fungsi init di dalam package main")
}

func main() {
    fmt.Println("Ini adalah fungsi main")
}

jika ingin menjalankan init function di dalam sebuah package tanpa menggunakan function lain di dalam nya, gunakan blank identifier
contoh 
import _"dasar/helper"

error
pada peraktinya panic digunakan untuk kondisi yang sangat fatal dan tidak dapat diperbaiki dan akan langsung menhentikan program, sementara error digunakan untuk situasi yang lebih umum di mana kegagalan dapat diperbaiki atau ditangani secara terstruktur oleh program.
golang memiliki interfce untuk membuat error nama interfacenya adalah error 
type error interface{
   Error string
}
untuk menggunakanya harus import "errors" untuk menjalankna error.New()
contoh:
package main
import (
	"errors"
	"fmt"
)

// Contoh fungsi yang mengembalikan error
func myFunction() error { //penggunaan interface error sebagai return 
	// Logika yang mungkin menghasilkan error
	if /* kondisi yang memicu error */ {
		return errors.New("Pesan error") //membuat error
	}
	// Jika tidak ada error
	return nil
}
func main() {
	// Memanggil fungsi yang mungkin mengembalikan error
	err := myFunction() //err akan bertipe data error karena menerima function yang ber return error
	// Memeriksa apakah ada error
	if err != nil {
		// Menangani error
		fmt.Println("Error:", err)
		return
	}
	// Menangani jika tidak ada error
	fmt.Println("Operasi berhasil")
}

contoh custom error:
// Custom error struct
type MyError struct {
	Message string
}
// Implementasi metode Error() untuk memenuhi interface error
func (e *MyError) Error() string {
	return e.Message
}
func main() {
	// Membuat instance custom error
	// "instance" mengacu pada penciptaan variabel dari suatu tipe data, baik itu tipe data
	err := &MyError{"Ini adalah custom error"}
fmt.Errorf()
	// Menangkap dan mencetak error
	fmt.Println("Error:", err)
}
/* kita  tidak perlu mengimpor package errors karena Anda tidak menggunakan errors.New()
kita hanya memanfaatkan kontrak interface error saja (sudah builtin seperti alias data bawaan lainnya)
dengan custom error kita bisa menkostumisasi data apa saja yang ingin di tampilkan*/

label
Dalam Go, label (label) digunakan untuk memberi nama pada loop for atau blok switch, sehingga Anda dapat menggunakan break atau continue bersama label untuk mengontrol aliran eksekusi program. Namun, label hanya berlaku untuk loop (for) atau blok switch yang diberi nama label tersebut. Penggunaan label memungkinkan Anda untuk mengatur aliran eksekusi dengan cara yang lebih spesifik dalam loop bersarang.
contoh:

loop:
   for{
    break loop  
   }