Konsep Dasar
- Context: Antarmuka untuk mengatur pembatalan, tenggat waktu, dan nilai-nilai yang terkait dengan suatu operasi.
- Context Hierarchy: Hierarki context memungkinkan pewarisan pembatalan dan nilai-nilai terkait.

Jenis Context
- Background Context: Context akar, digunakan sebagai parent untuk semua context lainnya. Diperoleh dari context.Background().tidak pernah di batalkan,tidak pernah timeout, tidak memiliki nilai.biasanya di miliki main function atau di dalam test atau awal proses request terjadi
- context.TODO(): Membuat context kosong yang digunakan sebagai placeholder jika context belum tersedia.(belum jelas context apa yang akan di gunakan)

context menganut konsep parent child
parent context bisa memiliki banyak child,tapi child context hanya bisa memiliki satu parent context
parent dant child context selalu terhubung

context merupakan data yang immutable,artinya setelah context di buat,di tidak bisa di ubah lagi.
ketika menambah value ke dalam context, atau menambah pengaturan timout dan lainnya. secara otomatis akan membentuk child context baru,bukan merubah context tersebut.
contoh:

    A
   ├── B
   │   ├── D
   │   └── E
   └── C
       ├── F
       └── G

jika parent context A di cancel, maka semua child dan sub child dari context A akan di batalkan
semua prilaku ini bersifat TURUNAN , ke atas nya tidak akan terpengaruh.
jika B di cancel semua yang ada di bawahnya akan tercancel, tetapi A tidak akan terpengaruh

Fungsi-fungsi Utama
- context.Background(): Membuat context kosong yang menjadi parent dari semua context lainnya.
- context.TODO(): Membuat context kosong yang digunakan sebagai placeholder jika context belum tersedia.
  
- context.WithCancel(parent Context) (ctx Context, cancelFn func()): Membuat context yang dapat dibatalkan dari context parent.ini di gunakan jika kita menjalankan proses lain dan kita ingin memberi sinyal cancle ke peroses tersebut.biasnya bentuknya berupa goroutine yang berbeda.gorutine yang menggunakan context harus melakukan pengecekan terhadap context nya.

- context.WithDeadline(parent Context, deadline time.Time) (ctx Context, cancelFn func()): Membuat context dengan tenggat waktu dari context parent.berbeda dengan timout yang kita tentukan adalaha durasinya, deadline yang kitat tentukan adalah spesifik waktunya , contohnya jika ingin kita cancel jam 12 siang
  
- context.WithTimeout(parent Context, timeout time.Duration) (ctx Context, cancelFn func()): Membuat context dengan tenggat waktu dari durasi waktu.selain menggunakan cancel, kita juga bisa mengirim sinya cancel secara otomatis menggunakan timeout
  
- context.WithValue(parent,key,value) : digunakan untuk membuat salinan context yang baru dengan nilai tambahan yang terkait dengan kunci tertentu.
  
Metode Context
- Deadline() (deadline time.Time, ok bool): Mendapatkan tenggat waktu dari context.
- Done() <-chan struct{}: Channel yang ditutup saat context dibatalkan atau mencapai tenggat waktu.
- Err() error: Mendapatkan kesalahan terkait dengan pembatalan atau tenggat waktu.
- Value(key interface{}) interface{}: Mendapatkan nilai terkait dengan kunci tertentu.jika value nya tidak ada , maka akan bertanya ke parent. tidak bisa tanya ke childnya