package testing adalah paket untuk melakukan pengetesan secara otomatis
testing.T : struct untuk melakukan unit test
testing.M : struct untuk mengatur life cycle(alur eksekusi tes)
testing.B : struct untuk benchmarking(mengukur kecepatan kode)

Aturan unit test
- nama file harus berakhiran "_test" jika nama filenya "hello.go" maka nama file testingnya "hello_testing.go"
- nama fungsi harus berawalan "Test" jika nama fungsinya "Say" maka nama fungsi testingnya "TestSay",untuk nama belakangnya tidak ada aturan harus sama yang penting depannya harus "Test"
- harus memiliki parameter (t *testing.T) dan tanpa mengembalikan nilai return

menjalan kan semua testing di directory
`go test`
verbose output
`go test -v`
menjalankan testing fungsi tertentu
`go test -v -run Testnamafungsi`
menjalan kan semua testing di sub directory (berguna jika running di root module)
`go test ./...`

menggagalkan unit test menggunakan panic bukan hal yang bagus, sebaiknya menggunakan 
fungsi Fail,FailNow,Error,dan Fatal yang sudah tersedia di testing.T
Fail : mengagalkan unit tes,tapi kode unit tes akan lanjut di eksekusi, di akhir tes akan menghasilkan gagal
FailNow : menggagal kan unit tes, eksekusi kode unit tes tidak akan di lanjutkan
Error : lebih seperti log(print), kemudian memanggil Fail
Fatal : seperti error tapi , yang kemudian di panggil FailNow

agar tidak melakukan pengecekan dengan menggunakan if secara beruntun, gunakan assertion
gunakan paket testify `go get -u github.com/stretchr/testify`, gunakan paket `github.com/stretchr/testify/assert`. di dalam modul testify ada 2 paket untuk melakukan assertion yaitu assert dan require
assert : jika gagal assert akan memanggil Fail
require : jika gagal require akan memanggil FailNow

membatalkan test(bukan menggagalkan)
bedanya membatalkan dan menggagalkan, menggalkan akan menghasilkan fail ketika testing, sedangkan membatalkan berarti unit test tidak akan di jalankan.contohnya jika kita punya fungsi yang hanya berjalan di linux dan kita ingin membuat test terhadap fungsi tersebut untuk tidak di test di windows. kita bisa membatalkannya menggunakana "skip"
t.Skip()

nama fungsinya harus "TestMain"
(m *testing.M)
m.Run() : menjalan kan semua unitest di dalam suatu paket
biasanya di gunakan untuk before after, melakukan sesuatu sebelum dan sesudah testing

nama fungsinya harus berawalan "Benchmark
(b *testing.B)
