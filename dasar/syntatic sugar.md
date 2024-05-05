Syntactic sugar adalah konsep dalam pemrograman yang merujuk pada fitur atau sintaks yang disediakan dalam bahasa pemrograman untuk membuat kode menjadi lebih mudah ditulis atau lebih mudah dibaca, tanpa menambahkan fungsionalitas baru ke bahasa tersebut. Dengan kata lain, itu adalah "gula sintaksis" yang membuat penulisan kode menjadi lebih "manis" atau lebih mudah dipahami bagi pengembang.

variable
- Pendeklarasian Variabel Tanpa Tipe dengan :=
x := 42 // Pendeklarasian variabel tanpa tipe
- Multiple Variable Assignment
var a, b, c = 1, 2, 3 // Deklarasi dan inisialisasi beberapa variabel dalam satu baris
- Blank Identifier 
_, err := someFunction() // Mengabaikan nilai kembalian yang tidak diperlukan

constanta
- Nilai Enumerasi dengan iota
const (
    Monday = iota // 0
    Tuesday      // 1
    Wednesday    // 2
)
- Multiple Constant Declarations
const (
    MinValue = 0
    MaxValue = 100
)
- Untyped Constants
const Pi = 3.14
- Expression-Based Constants
const (
    SecondsPerMinute = 60
    SecondsPerHour   = SecondsPerMinute * 60
)
- Named Constants
const HoursInDay = 24

type
- Membuat Alias
type ID string
type Score int
- Membuat Tipe Data Baru Berdasarkan Tipe Data yang Sudah Ada
type Celsius float64
- Membuat Tipe Data Struktur Baru
type Person struct {
    Name string
    Age  int
}
- Membuat Tipe Data Fungsi Baru
type BinaryFunc func(int, int) int
- Membuat Tipe Data Interfase Baru
type Writer interface {
    Write([]byte) (int, error)
}

array
- Inisialisasi Array
arr := [3]int{1, 2, 3} // Inisialisasi array dengan panjang 3
- Penugasan Nilai pada Array
arr[0] = 10
arr[1] = 20
- Inferensi Panjang Array
arr := [...]int{1, 2, 3} // Inferensi panjang array

if 
singkat
if x < 5 {
    fmt.Println("x kurang dari 5")
}
inisialisasi variable
if num := 10; num > 0 {
    fmt.Println("num adalah bilangan positif")
}
bersamaan return
if num := getValue(); num > 0 {
    fmt.Println("num adalah bilangan positif")
}

switch
tanpa kondisi
switch {
case x > 10:
    fmt.Println("x lebih dari 10")
case x < 0:
    fmt.Println("x kurang dari 0")
default:
    fmt.Println("x antara 0 dan 10")
}
fallthrough
switch grade {
case "A":
    fmt.Println("Pertahankan prestasi!")
    fallthrough
case "B":
    fmt.Println("Bagus!")
case "C":
    fmt.Println("Perlu ditingkatkan.")
}
kondisi bersamaan assigmen
switch age := 30; {
case age < 18:
    fmt.Println("Anda masih di bawah umur")
case age >= 18 && age < 65:
    fmt.Println("Anda dewasa")
default:
    fmt.Println("Anda seorang senior")
}

