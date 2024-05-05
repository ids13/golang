# fmt
| Verb       | Keterangan                                                                                                                                                                      |
| ---------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `%v`       | Menampilkan nilai apapun dengan representasi default.                                                                                                                           |
| `%T`       | Menampilkan tipe data dari variabel.                                                                                                                                            |
| `%d`       | Menampilkan bilangan bulat (integers) dalam format desimal.                                                                                                                     |
| `%b`       | Menampilkan bilangan bulat dalam format biner.                                                                                                                                  |
| `%x`, `%X` | Menampilkan bilangan bulat dalam format heksadesimal (lowercase atau uppercase).                                                                                                |
| `%o`       | Menampilkan bilangan bulat dalam format oktal.                                                                                                                                  |
| `%f`, `%F` | Menampilkan angka desimal dalam format titik-penambahan (floating point), dengan `%.nf` untuk menentukan jumlah digit desimal, di mana `n` adalah jumlah digit yang diinginkan. |
| `%e`, `%E` | Menampilkan angka desimal dalam format notasi ilmiah (scientific notation), dengan `%.nf` untuk menentukan jumlah digit desimal.                                                |
| `%s`       | Menampilkan string.                                                                                                                                                             |
| `%p`       | Menampilkan alamat pointer.                                                                                                                                                     |
| `%%`       | Menampilkan karakter persen secara harfiah.                                                                                                                                     |
| `%w`       | menampilkan error                                                                                                                                                               |
## ## Printf 
mencetak format teks dengan variabel yang diformat.
contoh
fmt.Printf("Format %s\n", variable)
## ## Println
mencetak teks diikuti dengan baris baru.
contoh
fmt.Println("Text")
## ## Sprintf
membuat string dari format teks dengan variabel yang diformat, tanpa mencetaknya ke output
## contoh
message = fmt.Sprintf("Format %s", variable)
//jadi hasil dari Sprintf tidak akan di tampilkan tapi akan langsung di tampung ke variable message
## Errorf
## membuat error baru dengan pesan yang diformat.
err = fmt.Errorf("Error message %d",var)

# Error
## ## New
## membuat error baru dengan pesan tertentu.
err = errors.New("Error message")
## ## Is
memeriksa apakah error merupakan instance dari kesalahan 
if errors.Is(err, specificError) {
    // Lakukan sesuatu
}
//specificError error di sini adalah hasil dari errors.new atau custom error yang di tampung ke dalam suatu variable error, dalam kasus di atas yaitu variable err
## ## Errorf
## membuat error baru dengan pesan yang diformat.
err = errors.Errorf("Error occurred
%d", errorCode)
bedanya dengan fmt.errorf,errorf menghasilkan error baru dengan format tertentu, sedangkan fmt.errorf menghasilkan error baru dengan format tertentu,dengan di balut string contohnya (fmt.errorf("keterangan
w",err)) untuk mengamil err nya saja,hasil dari fmt.errorf gunakan  errors.unwrap

# os
| Perizinan          | Angka Oktal | Simbol | Deskripsi                                |
| ------------------ | ----------- | ------ | ---------------------------------------- |
| Read-Write         | 6           | rw-    | Boleh membaca dan menulis                |
| Read-Execute       | 5           | r-x    | Boleh membaca dan mengeksekusi           |
| Write-Execute      | 3           | -wx    | Boleh menulis dan mengeksekusi           |
| Read-Write-Execute | 7           | rwx    | Boleh membaca, menulis, dan mengeksekusi |
| Read-Only          | 4           | r--    | Hanya boleh membaca                      |
| Write-Only         | 2           | -w-    | Hanya boleh menulis                      |
| Execute-Only       | 1           | --x    | Hanya boleh mengeksekusi                 |
| No Access          | 0           | ---    | Tidak ada akses                          |
## ## Getenv
## mendapatkan nilai dari variabel lingkungan (environment variable).
value = os.Getenv("VARIABLE_NAME")
## Args
## Slice yang berisi argumen baris perintah yang diberikan saat menjalankan program.
args = os.Args
## ## Exit
menghentikan program dan mengembalikan kode keluaran.
os.Exit(1)
## ## Mkdir
## membuat direktori baru.
err = os.Mkdir("directory_name", 0755)
## ## Remove
## menghapus file atau direktori.
err = os.Remove("file_or_directory_path")
## ## Rename
## mengubah nama file atau direktori.
err = os.Rename("old_name", "new_name")
## ## Open
## membuka file dalam mode tertentu.
file, err = os.Open("file_name")
## ## Create
## membuat atau membuka file baru untuk ditulis.
file, err = os.Create("file_name")
## ## Stat
## mendapatkan informasi tentang file atau direktori.
fileInfo, err = os.Stat("file_or_directory_path")
## ## Getwd
## mendapatkan direktori kerja saat ini.
dir, err = os.Getwd()
## ## Getuid
## mendapatkan ID pengguna (user ID) saat ini.
uid = os.Getuid()
## ## Hostname
## mendapatkan nama host sistem.
hostname, err = os.Hostname()

# flag
## Parse
mem-parsing argumen baris perintah dan mengisi nilai variabel yang ditentukan.
flag.Parse()
## String
## mendefinisikan dan mengatur nilai default untuk argumen bertipe string.
var strFlag = flag.String("name", "defaultValue", "Usage message")
## Int
## mendefinisikan dan mengatur nilai default untuk argumen bertipe integer.
var intFlag = flag.Int("num", defaultValue, "Usage message")
## Bool
## mendefinisikan dan mengatur nilai default untuk argumen bertipe boolean.
var boolFlag = flag.Bool("verbose", defaultValue, "Usage message")
## Args
mendapatkan daftar argumen posisional setelah argumen yang diproses(maksudnya data data yang tidak punya ## flag).
args = flag.Args()

# string
## Contains
## memeriksa apakah sebuah string mengandung substring tertentu.
contains = strings.Contains("hello world", "world")
## ToLower
## mengubah semua karakter dalam string menjadi huruf kecil.
lowercase = strings.ToLower("HELLO")
## ToUpper
## mengubah semua karakter dalam string menjadi huruf besar.
uppercase = strings.ToUpper("hello")
## TrimSpace
## menghapus spasi dari awal dan akhir string.
trimmed = strings.TrimSpace("  hello world  ")
## Split
## memisahkan string berdasarkan pemisah tertentu dan mengembalikan slice dari substring.
parts = strings.Split("hello world", " ")
## Join
## menggabungkan elemen-elemen slice menjadi satu string dengan pemisah tertentu.
joined = strings.Join([]string{"hello", "world"}, " ")
## Replace
## mengganti n kali kemunculan substring dalam string dengan substring lainnya.gunakan -1 untuk penggantian semua
replaced = strings.Replace("hello world", "world", "Go", 1)
## ReplaceAll
## mengganti semua kemunculan substring dalam string dengan substring lainnya
replaced = strings.Replace("hello world", "world", "Go")
## Index
## mencari indeks dari sebuah substring dalam string.
index = strings.Index("hello world", "world")
## Title
fmerubah setiap kata pertama dalam string menjadi huruf kapital
strings.Title("hello world")

strconv
## Atoi
string ke integer.
num, err = strconv.Atoi("123")
## Itoa
integer ke string.
str = strconv.Itoa(123)
## ParseFloat
string ke float dengan presisi tertentu.
f, err = strconv.ParseFloat("3.14", 64)
## ParseInt
string ke integer dengan basis tertentu.
num, err = strconv.ParseInt("1010", 2, 64) // Mengonversi dari basis 2 (binary) ke integer 64-bit
## ParseUint
string ke integer tak bertanda dengan basis tertentu.
num, err = strconv.ParseUint("FF", 16, 64) // Mengonversi dari basis 16 (hexadecimal) ke integer tak bertanda 64-bit
## FormatBool
boolean ke string.
str = strconv.FormatBool(true)
## FormatFloat
float ke string dengan presisi tertentu.
str = strconv.FormatFloat(3.14, 'f', -1, 64) // Presisi -1 untuk nilai default
## FormatInt
integer ke string dengan basis tertentu.
str = strconv.FormatInt(42, 10) // Basis 10 (desimal)
## FormatUint
integer tak bertanda ke string dengan basis tertentu.
str = strconv.FormatUint(255, 16) // Basis 16 (hexadecimal)
## AppendBool
representasi boolean ke slice byte.
## b = []byte("bool
.AppendBool(b, true)
## AppendFloat
float ke slice byte dengan presisi tertentu.
## b = []byte("float
.AppendFloat(b, 3.14, 'f', -1, 64)
## AppendInt
integer ke slice byte dengan basis tertentu.
## b = []byte("int
.AppendInt(b, 42, 10)
## AppendUint
integer tak bertanda ke slice byte dengan basis tertentu.
## b = []byte("uint
.AppendUint(b, 255, 16)

# math
## Abs
nilai absolut dari bilangan.
result = math.Abs(-10.5)
## Ceil
nilai terkecil yang lebih besar dari atau sama dengan bilangan desimal.
result = math.Ceil(10.3)
## Floor
nilai terbesar yang kurang dari atau sama dengan bilangan desimal.
result = math.Floor(10.8)
## Round
langan desimal ke bilangan terdekat.
result = math.Round(10.49)
## Sqrt
akar kuadrat dari bilangan.
result = math.Sqrt(16)
## Min
nilai terkecil dari dua bilangan.
result = math.Min(10, 20)
## Max
nilai terbesar dari dua bilangan.
result = math.Max(10, 20)
## Pow
pangkat dari suatu bilangan.
result = math.Pow(2, 3)

# container/list
## Initiating List:
import "container/list"
// Create a new list
myList := list.New()
## Adding Elements:
// Add element to the back of the list
myList.PushBack(10)
// Add element to the front of the list
myList.PushFront(20)
## Accessing Elements:
// Get the first element
firstElement := myList.Front().Value
// Get the last element
lastElement := myList.Back().Value
## Iterating Through List:
// Iterate forward
for element := myList.Front(); element != nil; element = element.Next() {
    // Access element.Value
}
// Iterate backward
for element := myList.Back(); element != nil; element = element.Prev() {
    // Access element.Value
}
## Removing Elements:
// Remove element from the front
firstElement := myList.Front()
myList.Remove(firstElement)
// Remove element from the back
lastElement := myList.Back()
myList.Remove(lastElement)
## Checking Length:
// Get the length of the list
length := myList.Len()
## Clearing List:
// Clear the list
myList.Init()

# container/ring
## Inisialisasi Ring:
Membuat ring kosong dengan panjang tertentu.
import "container/ring"

r := ring.New(5) // Membuat ring dengan panjang 5
## Menambahkan Nilai:
Menambahkan nilai ke dalam ring pada posisi saat ini.
r.Value = 1
r = r.Next() // Pindah ke node berikutnya
## Iterasi Ring: 
Melakukan iterasi melalui semua elemen dalam ring.
r.Do(func(p interface{}) {
    // Lakukan sesuatu dengan setiap elemen
    fmt.Println(p)
})
## Mengakses Nilai: 
Mengakses nilai dari posisi saat ini dalam ring.
fmt.Println(r.Value)
## Pindah ke Node Berikutnya: 
Memindahkan pointer ke node berikutnya dalam ring.
r = r.Next()
## Pindah ke Node Sebelumnya: 
Memindahkan pointer ke node sebelumnya dalam ring.
r = r.Prev()
## Panjang Ring: 
Mendapatkan panjang ring.
length := r.Len()

# sort
## mengurutkan slice
sort.Ints(a)
sort.Float64s(a)
sort.Strings(a)
## Pencarian di slice
sort.SearchInts(a, x) int
sort.SearchFloat64s(a, x) int
sort.SearchStrings(a, x) int
## pengurutan custom slice
sort.Slice(data, func(i, j int) bool { // logika pengurutan kustom })
contoh :
	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println("Sorted by Age:", people)

## pengurutan custom slice dengan interface
type Data []int // Ganti int dengan tipe data sesuai kebutuhan(tipe data daras atau struct)
// Implementasikan antarmuka sort.Interface untuk Data
func (d Data) Len() int           { return len(d) }
func (d Data) Less(i, j int) bool { return d[i] < d[j] } //jika menggunakan struct slice tunjuk.data mana yang di bandingkan
func (d Data) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

contoh :
// Definisikan tipe data custom
type Person struct {
    Name string
    Age  int
}
// Definisikan tipe slice untuk Person
type People []Person
// Implementasikan antarmuka sort.Interface untuk People
func (p People) Len() int           { return len(p) }
func (p People) Less(i, j int) bool { return p[i].Age < p[j].Age }
func (p People) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func main() {
    // Inisialisasi slice data
    people := People{
        {"Alice", 30},
        {"Bob", 25},
        {"Charlie", 35},
    }
    // Lakukan pengurutan menggunakan sort.Sort
    sort.Sort(people)
    // Cetak hasil pengurutan
    fmt.Println(people)
}

# time
## Membuat Waktu:
time.Now(): Membuat waktu saat ini.
time.Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location): Membuat waktu dengan nilai yang ditentukan.
## Manipulasi Waktu:
t.Add(d Duration) Time: Menambahkan durasi ke waktu t.
t.Sub(u Time) Duration: Menghitung durasi antara dua waktu.
t.Format(layout string) string: Mengonversi waktu menjadi string dengan format tertentu.
t.Parse(layout string) (timer.time,err) : mengonversi string menjadi format waktu
## Mengambil Komponen Waktu:
t.Year() int, t.Month() Month, t.Day() int: Mengambil tahun, bulan, dan hari dari waktu t.
t.Hour() int, t.Minute() int, t.Second() int: Mengambil jam, menit, dan detik dari waktu t.
## Pembandingan Waktu:
t.Before(u Time) bool, t.After(u Time) bool: Memeriksa apakah waktu t sebelum atau setelah waktu u.
t.Equal(u Time) bool: Memeriksa apakah waktu t sama dengan waktu u.
## Zona Waktu:
time.LoadLocation(name string) (*Location, error): Memuat lokasi dari zona waktu dengan nama tertentu.
time.FixedZone(name string, offset int) *Location: Membuat zona waktu tetap dengan offset yang ditentukan.offset di hijtung dalam satuan detik jadi jika butuh 2 jam = 60*60*2
## Konstanta Waktu:
time.Second, time.Minute, time.Hour: Konstanta untuk durasi satu detik, satu menit, dan satu jam.
time.RFC3339, time.RFC822: Konstanta format untuk RFC3339 (ISO 8601) dan RFC822 (RFC 822/1123) format.
## Durasi:
time.Duration: Representasi durasi waktu, digunakan untuk menambah atau mengurangkan waktu.
	
/*
	   layout waktu di golang

	   time.RFC3339: "2006-01-02T15:04:05Z07:00" (ISO 8601)
	   time.RFC822: "02 Jan 06 15:04 MST"
	   time.RFC822Z: "02 Jan 06 15:04 -0700"
	   time.RFC850: "Monday, 02-Jan-06 15:04:05 MST"
	   time.RFC1123: "Mon, 02 Jan 2006 15:04:05 MST"
	   time.RFC1123Z: "Mon, 02 Jan 2006 15:04:05 -0700"

	   2006	layout untuk tahun
	   01 		layout untuk bulan
	   02 		layout untuk tanggal
	   03/15 	layout untuk jam 12jam/24jam
	   04 		layout untuk menit
	   05 		layout untuk detik
	   PM/AM 	layout tanda waktu

	   Format Tanggal:
	   "2006-01-02"
	   "02 Jan 2006"
	   "January 02, 2006"
	   "02/01/2006"
	   "Jan-02-2006"
	   Format Waktu:
	   "15:04:05"
	   "03:04 PM"
	   "03:04:05 PM"
	   "15:04"

	*/
# reflect
## Mengambil Informasi Tipe Data:
reflect.TypeOf(i interface{}) reflect.Type: Mengembalikan tipe dari nilai yang diberikan.
reflect.ValueOf(i interface{}) reflect.Value: Mengembalikan nilai dari sebuah tipe.
contoh : 
t := reflect.TypeOf(x)  // Mengambil tipe dari nilai x
v := reflect.ValueOf(x) // Mengambil nilai dari variabel x
## Memeriksa Tipe dan Nilai:
t.Kind() reflect.Kind: Mengembalikan jenis tipe nilai.
v.Interface() interface{}: Mengembalikan nilai dalam bentuk interface{}.
## Mendapatkan Informasi Struktur:
t.NumField() int: Mengembalikan jumlah field pada tipe struktur.
t.Field(i int) reflect.StructField: Mengembalikan informasi tentang field ke-i dari tipe struktur.
## Tag:
field.Tag: Mendapatkan tag yang terkait dengan field.
## Mendapatkan Informasi Tambahan:
t.NumMethod() int: Mengembalikan jumlah metode yang dimiliki oleh tipe.
## Manipulasi Nilai:
v.CanSet() bool: Memeriksa apakah nilai dapat diubah.
v.FieldByName(name string) reflect.Value: Mengembalikan nilai atribut dengan nama tertentu.
v.FieldByName(name string).Set(value reflect.Value): Mengatur nilai atribut dengan nama tertentu.
## Konversi Tipe:
v.Convert(t reflect.Type) reflect.Value: Mengubah nilai menjadi tipe yang diinginkan.

# regex
| Pola Regex | Deskripsi                                 | Contoh Pola | Contoh Cocok            |
| ---------- | ----------------------------------------- | ----------- | ----------------------- |
| `.`        | Satu karakter apa pun kecuali newline.    | `a.b`       | `acb`, `a!b`            |
| `\w`       | Karakter alfanumerik atau underscore.     | `\w+`       | `hello_world`, `abc123` |
| `\W`       | Bukan karakter alfanumerik.               | `\W+`       | `!@#$`, `   `           |
| `\d`       | Digit (0-9).                              | `\d+`       | `123`, `007`            |
| `\D`       | Bukan digit.                              | `\D+`       | `abc`, `!@#`            |
| `\s`       | Whitespace (spasi, tab, newline).         | `\s+`       | ` `, `\t`, `\n`         |
| `\S`       | Bukan whitespace.                         | `\S+`       | `hello`, `123abc`       |
| `*`        | 0 atau lebih kemunculan.                  | `ab*c`      | `ac`, `abc`, `abbc`     |
| `+`        | 1 atau lebih kemunculan.                  | `ab+c`      | `abc`, `abbc`           |
| `?`        | 0 atau 1 kemunculan.                      | `ab?c`      | `ac`, `abc`             |
| `{n}`      | Tepat n kemunculan.                       | `a{3}`      | `aaa`, `aaab`           |
| `{n,}`     | Setidaknya n kemunculan.                  | `a{2,}`     | `aa`, `aaa`             |
| `{n,m}`    | Antara n dan m kemunculan.                | `a{2,4}`    | `aa`, `aaa`, `aaaa`     |
| `^`        | Awal string.                              | `^abc`      | `abcdef`, `abc`         |
| `$`        | Akhir string.                             | `abc$`      | `123abc`, `defabc`      |
| `\b`       | Batas kata.                               | `\bword\b`  | `word`, `password`      |
| `[abc]`    | Salah satu karakter dalam kurung.         | `[abc]`     | `a`, `b`, `c`           |
| `[^abc]`   | Karakter yang tidak ada dalam kurung.     | `[^abc]`    | `d`, `e`, `f`           |
| `[a-z]`    | Karakter dalam rentang a-z.               | `[a-z]`     | `a`, `b`, ..., `z`      |
| `[A-Z]`    | Karakter dalam rentang A-Z.               | `[A-Z]`     | `A`, `B`, ..., `Z`      |
| `(abc)`    | Membuat grup dari pola "abc".             | `(abc)`     | `abc`, `abcdef`         |
| `(?:abc)`  | Grup non-capturing dari pola "abc".       | `(?:abc)`   | N/A                     |
| `(?=abc)`  | Posisi yang diikuti oleh "abc".           | `(?=abc)`   | `xyzabc`, `abc123`      |
| `(?!abc)`  | Posisi yang tidak diikuti oleh "abc".     | `(?!abc)`   | `xyz`, `123`            |
| `\|`       | Operasi OR, cocok dengan salah satu pola. | `a\|b`      | `a`, `b`                |
| `\.`       | Cocok dengan titik.                       | `\.`        | `.`                     |
| `\n`       | Cocok dengan karakter newline.            | `\n`        | `\n`                    |



## Compile Pattern:
regexp.MustCompile(pattern string) *Regexp: Membuat objek Regexp dari pola regex yang diberikan. Panik jika ada kesalahan sintaksis.
## Compile Pattern (dengan Penanganan Error):
regexp.Compile(pattern string) (*Regexp, error): Membuat objek Regexp dari pola regex yang diberikan. Mengembalikan error jika ada kesalahan sintaksis.
## Match String:
re.MatchString(s string) bool: Memeriksa apakah pola regex cocok dengan string yang diberikan.
## Find All Matches:
re.FindAllString(s string, n int) []string: Mencari semua kemunculan pola regex dalam string dan mengembalikan array dari string yang cocok.
## Find First Match:
re.FindString(s string) string: Mencari kemunculan pertama pola regex dalam string dan mengembalikan string yang cocok.
## Replace Matches:
re.ReplaceAllString(src, repl string) string: Mengganti semua kemunculan pola regex dalam string dengan string pengganti.

# encode 
## encoding/json
Marshal(v interface{}) ([]byte, error): Mengonversi tipe data menjadi JSON.
Unmarshal(data []byte, v interface{}) error: Mengonversi JSON menjadi tipe data.
## encoding/xml
Marshal(v interface{}) ([]byte, error): Mengonversi tipe data menjadi XML.
Unmarshal(data []byte, v interface{}) error: Mengonversi XML menjadi tipe data.
## encoding/csv
NewWriter(w io.Writer) *Writer: Membuat penulis CSV baru.
NewReader(r io.Reader) *Reader: Membuat pembaca CSV baru.
WriteAll(records [][]string) error: Menulis semua rekaman ke penulis CSV.
ReadAll() (records [][]string, err error): Membaca semua rekaman dari pembaca CSV.
## encoding/base64
EncodeToString(src []byte) string: Mengonversi slice byte menjadi string base64.
DecodeString(s string) ([]byte, error): Mengonversi string base64 menjadi slice byte.
## encoding/hex
EncodeToString(src []byte) string: Mengonversi slice byte menjadi string heksadesimal.
DecodeString(s string) ([]byte, error): Mengonversi string heksadesimal menjadi slice byte.
## net/url
QueryEscape(s string) string: Mengonversi string menjadi bentuk yang dapat digunakan dalam URL.
QueryUnescape(s string) (string, error): Mengonversi bentuk yang dapat digunakan dalam URL menjadi string.

# bufio
## Membuat Scanner Baru:
NewScanner:
Fungsi: Membuat scanner baru untuk membaca input dari io.Reader.
Contoh Penggunaan: scanner := bufio.NewScanner(reader)
Membaca Input dengan Scanner:
Scan:
Fungsi: Membaca baris berikutnya dari input.
Contoh Penggunaan: scanner.Scan()
Text:
Fungsi: Mengembalikan teks dari baris terakhir yang dipindai.
Contoh Penggunaan: text := scanner.Text()
## Membuat Reader Baru dengan Buffering:
NewReader:
Fungsi: Membuat reader baru yang melakukan buffering untuk io.Reader.
Contoh Penggunaan: reader := bufio.NewReader(reader)
## Membaca Input dengan Reader:
ReadBytes:
Fungsi: Membaca data dari io.Reader hingga menemui delimiter tertentu.
Contoh Penggunaan: data, err := reader.ReadBytes('\n')
ReadString:
Fungsi: Membaca string dari io.Reader hingga menemui delimiter tertentu.
Contoh Penggunaan: line, err := reader.ReadString('\n')
## Menulis Output dengan Buffering:
NewWriter:
Fungsi: Membuat writer baru untuk menulis output ke io.Writer dengan buffering.
Contoh Penggunaan: writer := bufio.NewWriter(writer)
Flush:
Fungsi: Memastikan bahwa semua data ditulis ke io.Writer.
Contoh Penggunaan: writer.Flush()


# file
| Fungsi atau Metode | Deskripsi                                                                                             |
| ------------------ | ----------------------------------------------------------------------------------------------------- |
| os.Create()        | Membuat file baru atau membuka file jika sudah ada untuk ditulis.                                     |
| os.Open()          | Membuka file yang sudah ada untuk dibaca.                                                             |
| os.OpenFile()      | Membuka file dengan kontrol lebih besar terhadap mode akses, flag, dan permission.                    |
| os.ReadFile()      | Membaca isi dari file secara keseluruhan dan mengembalikan slice byte.                                |
| os.WriteFile()     | Menulis data ke file.                                                                                 |
| os.Remove()        | Menghapus file atau symlink.                                                                          |
| os.Mkdir()         | Membuat direktori baru.                                                                               |
| os.ReadDir()       | Membaca isi dari direktori dan mengembalikan slice os.DirEntry yang mewakili entri-entri di dalamnya. |
| os.Rename()        | Mengubah nama atau memindahkan file atau direktori.                                                   |
| os.Stat()          | Memberikan informasi tentang file atau direktori.                                                     |
| os.Chmod()         | Mengubah permission file atau direktori.                                                              |
| os.Getwd()         | Mendapatkan direktori kerja saat ini.                                                                 |

## mode
| Mode        | Deskripsi                                                                           |
| ----------- | ----------------------------------------------------------------------------------- |
| os.O_RDONLY | File dibuka hanya untuk operasi baca saja.                                          |
| os.O_WRONLY | File dibuka hanya untuk operasi tulis saja.                                         |
| os.O_RDWR   | File dibuka untuk operasi baca dan tulis.                                           |
| os.O_CREATE | Jika file tidak ada, file akan dibuat.                                              |
| os.O_APPEND | Data akan ditulis di akhir file.                                                    |
| os.O_TRUNC  | File akan dipotong menjadi kosong jika file sudah ada.(overwrite)                   |
| os.O_EXCL   | Jika file sudah ada, os.OpenFile() akan gagal dengan error os.ErrExist.             |
| os.O_SYNC   | File akan di-synchronize (dijaga konsistensinya) setiap kali ada penulisan ke file. |

gunakan opratror `|` untuk mengkombinsasikan mode.


