goroutine simple nya adalah thread ringan. thread berjalan secara pararel, goroutine berjalan secara concuren
goroutine 2kb sedangkan thread 1mb,goroutine berjalan di atas thread.goroutine bisa berjalan secara pararel ataupun concuren tergantung jumlah core dan penjadwalan runtime
Concurrency (Konkurensi): Berfokus pada bagaimana program menangani banyak tugas secara bersamaan. Dalam konkurensi, tugas-tugas dapat dimulai, dijalankan, dan selesai secara bersamaan, tetapi tidak menjamin bahwa mereka akan dieksekusi secara simultan di beberapa core CPU.
Parallelism (Paralelisme): Berfokus pada eksekusi beberapa tugas secara bersamaan secara simultan di beberapa core CPU yang berbeda. Dalam paralelisme, tugas-tugas dijalankan secara benar-benar simultan pada core CPU yang berbeda, meningkatkan kinerja secara langsung.
pararel :bersamaan
concuren :bergantian
di pararel butuh banyak thread sedangkan di concuren sebaliknya

cpu-bound
algoritma yang yang mengandalkan kecepatan cpu
contohnya machine learning
kurang cocok jika menggunakan concuren, akan lebih terbantu jika menggunakan pararel

i/o bound
algoritma yang bergantung pada kecepatan input output device yang di gunakan
contohnya membaca data dari file,data base, dll

Goroutine adalah cara di Go untuk melakukan konkurensi. Mereka memungkinkan Anda untuk menjalankan fungsi secara independen secara bersamaan dengan kode lainnya. Goroutine berbagi waktu CPU dan berjalan secara bersamaan dengan goroutine lainnya dalam satu proses. Namun, tidak ada jaminan bahwa goroutine akan dieksekusi secara simultan di beberapa core CPU; penjadwalannya ditangani oleh runtime Go. Sehingga, walaupun goroutine memungkinkan eksekusi bersamaan, mereka bukanlah representasi langsung dari paralelisme. Paralelisme dapat dicapai dengan menjalankan beberapa goroutine secara bersamaan di beberapa core CPU yang berbeda.

goroutine di jalankan goscheduler dalam thread,dimana jumlah threadnya sebanyak GOMAXPROCS(biasanya jumlah core CPU)
didalam go scheduler terdapat 3 istilah
G : Goroutine
M : thread (Machine)
P : Processor
-untuk membuat goroutine tinggal tambahkan kata "go" sebelum memanggil fungsi yang akan di jalankan dalam goroutine
-sebuah fungsi yang di jalankan dengan goroutine akan di jalankan secaray asycronous artinya tidak akan 
di tunggu sampa fungsi tersebut selesai.aplikasi akan lanjut berjalan ke kode selanjutnya tanpa menunggu
fungsi yang dijalankan secara goroutine selesai.
-tidak cocok menjalankan goroutine di fungsi yang memiliki nilai return.

channel
-channel adalah tempat komunikasi secara sychronous bisa di lakukan goroutine
-di channel ada pengirim dan penerima, biasanya pengirim dan penerima adalah goroutine yang berbeda
-saat melakukan pengiriman data,goroutine akan terblok sampai ada yang menerima data tersebut
maka dari itu channel di sebut sychronous(blocking)
-channel cocok sekali sebagai sebagai alternatif async await (seperti di bahasa pemrogramman lainnya)
sifat
-secara default channel hanya menampung satu data.jika ingin mengirim data lagi , harus menunggu data di channel di ambil
-channel hanya menerima satu tipe data
-channel bisa di ambil lebih dari satu goroutine
-harus di close jika tidak di gunakan, jika tidak akan menyebabkan memory leak.
saran gunakan defer untuk close, karena defer meskipunt terjadi error , kode yang di defer akan tetap dijalankan

channel di representasikan dengan tipe data "chan"
bisa di buat menggunakan make()
saat pembuatan harus di tentukan tipe data yang akan di masukan ke dalam channel
contoh :
channel := make(chan string)

mengirim : channel <- data
menerima : data <- channel
menutup : close(channel)
<-channel : mengambil data dari channel.(tanpa di masukan ke variable)

jika mengirim tanpa ada yang menerima akan error "panic send on close channel"
jikat menerima tanpa ada yang kirim akan error "deadlock"

sebagai parameter
di golang parameter secara default pass by value , jika ingin pass by reference harus menggunakan pointer.
sedangkan channel tidak perlu melakukan hal tersebut jika ingin menggunakan channel sebagai parameter
contoh:

func Hello(channel chan string){}
func main(){
    c := make(chan string)
    go Hello(c)
}

saat mengirim channel sebagai parameter,channel tersebut bisa di gunakan untuk mengirim dan menerima,
kita bisa juga menandai channel tersebut di gunakan untuk mengirim atau menerima
mengirim(memasukan data ke dalam channel) : func (channel chan <- string)
menerima(mengambil data dari channel )    : func (channel <- chan string)

secara default channel hanya bisa menampung satu data, jadi jika kita ingin menambahkan data ke dua maka harus menunggu sampai data ke 1 di ambil terlebih dahulu,kadang ada kasus di mana pengirim lebih cepat di bandingkan penerima,jika menggunakan channel saja maka pengirim akan menjadi lambat.
buffered channel gunanya untuk menampung data antrian di channel
caranya :
channel := make(chan string,3)
dalam kasus ini kita membuat tiga buffer.jika kita menambahkan buffer, meskipun data nya tidak ada yang mengambil tidak akan terjadi error.
len(channel) : melihat jumlah data di buffer
cap(channel) : melihat panjang buffer

kadang ada kasus sebuah channel di kirim data terus menerus oleh pengirim.dan tidak jelas kapan channel tersebut berhenti menerima data.salah satu cara menanganinya adalah dengan range,ketika sebuah channel di close maka perulangannya akan berhenti.jadi pastikan di close channel nya atau akan menghasilkan deadlock

membuat beberapa channel dan menjalankan goroutine secara bersaamaan dan kita ingin mendapatkana data dari semua channel maka kita bisa menggunakan select.dengan select channel kita bisa memilih data tercepat dari beberapa channel, jika data datang secara bersamaan maka akan di pilih random.jika default di setting maka selama belum ada data yang masuk ke channel, kode di dalam default akan di eksekusi terus menerus

mutex(mutual exclusion),hanya satu goroutine yang dapat mengakses data pada satu waktu,biasa di gunakan untuk mengunci sebuah kode program.

sync.RWMutex (Read write mutex),terkadang ada kasus di mana kita ingin melakukan locking tidak hanya ketika merubah data , tapi juga ketika membaca,data.sebenarnya bisa saja menggunakan mutex biasa tapi nanti akan rebutan antara proses rubah dan baca(rececondition).

var RWMutex sync.RWMutex
//ini untuk write(edit)
RWMutex.Lock()
RWMutex.Unlock()
//ini untuk readnya
RWMutex.RLock()
RWMutex.RUnlock()

Mutex:
-sync.Mutex memungkinkan hanya satu goroutine pada satu waktu untuk mengunci sumber daya.
Ini bersifat eksklusif, artinya ketika satu goroutine telah mengunci mutex, tidak ada goroutine lain yang dapat mengakses sumber daya yang sama sampai mutex dilepaskan.

RWMutex:
-sync.RWMutex memungkinkan banyak goroutine untuk membaca sumber daya secara bersamaan, tetapi hanya satu goroutine yang dapat menulis pada satu waktu.
-Ini memungkinkan kunci baca (RLock) yang bisa dilakukan secara bersamaan oleh beberapa goroutine, tetapi hanya memungkinkan satu kunci tulis (Lock) pada satu waktu.

deadlock adalah kondisi di mana sebuah prosess go routine saling menunggu lock.

sync.WaitGroup di gunakan untuk menunggu sejumlah runtime(goroutine) selesai di eksekusi sebelum melanjutkan eksekusi program utama
method:
add(jumlah): untuk menandai bahwa ada proses goroutine(waitgroup.add(1) berarti menambahkan satu kedalam waitgroup,di contoh, ketika di dalam for loop kita add(1)berarti menambahkan satu secara looping ke waitgroup, dan total di waitgroup yang akan di gunakann oleh waitgroup.wait())
done(): untuk menandai bahwa proses goroutine sudah selesai(biasanya bersamaan dengan defer), ini biasanya di letakan sebelum goroutine di jalankan
wait(): untuk menunggu proses selesai

sync.once
memastikan sebuah fungsi di eksekusi hanya satu kali,jadi berapapun goroutine yang mengakses, hanya goroutine pertama yang akan mengeksekusi.goroutine yg lain akan di abaikan

sync.Pool adalah struktur data di Go yang digunakan untuk menyimpan objek-objek yang bisa digunakan kembali (pooled). Tujuan utama dari sync.Pool adalah untuk mengurangi alokasi memori baru dengan memanfaatkan kembali objek-objek yang sudah ada sebelumnya(aman dari race condition).kita bisa mengambil dan menyimpan data di pool
Fungsi New dalam sync.Pool adalah sebuah field yang menentukan fungsi atau closure yang akan dipanggil saat pool perlu membuat objek baru. Ini memungkinkan pool untuk secara dinamis membuat objek baru sesuai dengan kebutuhan, misalnya ketika pool kosong dan diperlukan objek baru.
-deklarasi
var pool = sync.Pool{} 
-menyimpan data ke pool
pool.put(data)
-mengambil data dari pool
data = pool.get()

sync.map fungsinya sama seperti map di golang, bedanya ini aman untuk concuren goroutine

// Inisialisasi sync.Map
var sm sync.Map
// Menambahkan elemen
sm.Store(key, value)
// Mengakses elemen
value, ok := sm.Load(key)
// Menghapus elemen
sm.Delete(key)
// Iterasi melalui elemen
sm.Range(func(key, value interface{}) bool {
    // Lakukan sesuatu dengan key dan value
    return true // untuk melanjutkan iterasi, false untuk berhenti
})

sync.cond adalah locker berbasis kondisi,
wait() : menunggu
signal() : memberitahu sebuah goroutine untuk jalan
broadcast() memberitahu semua goroutine untuk jalan
untuk membuat cond, menggunakan fungsi sync.NewCond(locker)
sebelum menjalankan wait() cond harus di lock terlebiih dahulu, dan sesudahnya harus di unlock

atomic adalah bagian dari standar library Go (sync/atomic) yang menyediakan operasi atomik untuk mengakses dan memodifikasi variabel berbasis tipe data primitif.
-Load dan Store: Operasi ini digunakan untuk membaca dan menulis nilai variabel atomik.
sebaiknya dipertimbangkan dengan hati-hati. Dalam beberapa kasus, penggunaan struktur data yang lebih kompleks dalam atomic.Value dapat menghasilkan overhead dan kompleksitas tambahan yang tidak diinginkan.
-Add dan Sub: Operasi ini digunakan untuk menambahkan atau mengurangkan nilai variabel atomik dengan nilai tertentu.
-CompareAndSwap (CAS): Operasi ini membandingkan nilai variabel atomik dengan nilai yang diberikan. Jika nilai variabel sama dengan nilai yang dibandingkan, maka nilai variabel akan diganti dengan nilai baru. Operasi ini berguna untuk mengimplementasikan mekanisme penguncian (lock-free) di beberapa kasus.
-Swap: Operasi ini menggantikan nilai variabel atomik dengan nilai baru dan mengembalikan nilai lama.
-Increment dan Decrement: Operasi ini digunakan untuk menambahkan atau mengurangkan nilai variabel atomik dengan nilai 1.

Timer digunakan untuk menjadwalkan satu atau beberapa tugas untuk dijalankan di masa mendatang setelah jangka waktu tertentu. Ini berguna saat Anda perlu menunda eksekusi tugas dalam program Anda.ketika waktunya expire, event akan di kirim ke channel.

-untuk membuat timer 
timmer = timer.NewTimer(duration)
-menerima sinyal dari timer
<-timer.C
-mematikan timer
timer.Stop()

time.after()
kadang kita butuh channel nya saja , tidak butuh data timmernya

time.afterFunction()
menjalankan sebuah fungsi dengan delay waktu tertentu
tidak perlu menggunakan channelnya,cukup kirim fungsi yang akan di panggil ketika timmer mengirim kejadian

time.ticker digunakan untuk membuat pengulangan periodik (interval) sesuai dengan waktu yang ditentukan.ketika waktu ticker sudah expire,maka event akan di kirim ke channel
-membuat ticker
ticker := time.NewTicker(time.Second)
-menerima event ticker
case <-ticker.C
-menghentikan ticker
ticker.Stop()

time.tick(duration)
hanya butuh channelnya saja, tidak butuh data timernya

GOMAXPROCS
goroutine sebenarnya berjalan di dalam thread, untuk mengetahu jumlah thread yang berjalan ketika aplikasi di jalankan kita bisa menggunakan fungsi GOMAXPROCS yang ada di dalam package runtime.
secara default, jumalah thread di golang sesuai dengan jumlah core cpu 
kita juga bisa melihat jumalh core cpu denga menggunakan fungsi runtime.NumCpu()
