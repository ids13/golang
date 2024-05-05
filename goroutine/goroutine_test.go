package goroutine

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func RunHello() {
	println("hello world")
}

func DisplayNumber(n int) {
	println("number", n)
}

func ParamChan(channel chan string) {
	channel <- "hallo"
}
func TestHello(t *testing.T) {
	go RunHello()
	println("halo dunia")
	//tambahkan sleep karena program akan berjalan ke perintah berikutnya, tidak akan menunggu goroutinenya selesai
	time.Sleep(5 * time.Second)
}

func InChan(channel chan<- string) {
	channel <- "dari IN"
	time.Sleep(2 * time.Second)
}

func OutChan(channel <-chan string) {
	data := <-channel
	println(data)
}

func TestMany(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go func() {
		channel <- "hallo"
		println("data terkirim ke channel")
	}()
	//menunggu data di kirim kan dulu ke channel
	time.Sleep(2 * time.Second)
	data := <-channel
	println(data)
	//menunggu sampai semuanya selesai di eksekusi
	time.Sleep(5 * time.Second)
}

func TestParamChan(t *testing.T) {
	c := make(chan string)
	go ParamChan(c)
	time.Sleep(2 * time.Second)
	data := <-c
	println(data)
	time.Sleep(2 * time.Second)
}

func TestInOutChan(t *testing.T) {
	c := make(chan string)
	defer close(c)
	go InChan(c)
	go OutChan(c)
	time.Sleep(5 * time.Second)
}

func TestBuff(t *testing.T) {
	c := make(chan string, 3)
	defer close(c)
	println(len(c))
	println(cap(c))
	c <- "test 1"
	c <- "test 2"
	c <- "test 3"

	println(<-c)
	println(<-c)
	println(<-c)
}

func TestRange(t *testing.T) {
	c := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			c <- "test range " + strconv.Itoa(i)
		}
		//	close(c)
	}()

	for val := range c {
		fmt.Println(val)
	}
}

func TestSelect(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)
	go ParamChan(c1)
	go ParamChan(c2)

	counter := 0
	// harus menggunakan infinite looping karena jika tidak selectcase akan di jalankan hanya sekali
	for {
		select {
		case data := <-c1:
			fmt.Println("data dari channel 1 = ", data)
			counter++
		case data := <-c2:
			fmt.Println("data dari channel 1 = ", data)
			counter++
		//jika datanya belum masuk maka kode di default akan di eksekusi terus menerus
		default:
			fmt.Println("menunggu data")
		}
		// jika looping tidak di hentikan maka looping akan berusaha mengambil data,
		// tapi tidak pernah ada data lagi yang masuk, jadi akan menimbulkan deadlock
		if counter == 2 {
			break
		}
	}
}

func work(done chan bool) {
	fmt.Println("Doing some work...")
	time.Sleep(time.Second) // Menunggu 1 detik (contoh pekerjaan)
	done <- true            // Mengirimkan nilai ke channel setelah pekerjaan selesai
}

func TestWork(t *testing.T) {
	done := make(chan bool)
	go work(done) // Memulai goroutine untuk melakukan pekerjaan
	go work(done)
	<-done // Mengambil nilai dari channel, tetapi tidak disimpan ke dalam variabel
	fmt.Println("Work is done!")
}

func TestRaceCondition(t *testing.T) {
	x := 0
	for i := 1; i <= 1000; i++ {
		/*kettika di jalankan hasilnya kurang dari 100.000(harusnya 100.000)
		ini terjadi karena racecondition,ada 1000 goroutine yang mengakse varible yang sama
		ada beberapa goroutine yang mengakses x dengan posisi nilai yang sama, jad
		*/
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("counter : ", x)

}

func TestMutex(t *testing.T) {
	var mutex sync.Mutex
	x := 0
	for i := 1; i <= 1000; i++ {

		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("counter : ", x)
}

type bankAccount struct {
	RWMutex sync.RWMutex
	balance int
}

func (a *bankAccount) AddBalance(mount int) {
	a.RWMutex.Lock()
	a.balance += mount
	a.RWMutex.Unlock()
}
func (a *bankAccount) GetBalance() int {
	a.RWMutex.RLock()
	data := a.balance
	a.RWMutex.RUnlock()
	return data
}

func TestRWMutex(t *testing.T) {
	account := bankAccount{}
	for i := 1; i <= 10; i++ {
		go func() {
			for j := 1; j <= 5; j++ {
				account.AddBalance(1)
				fmt.Println("balance = ", account.GetBalance(), " dari go ke ", i)
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("total balance = ", account.GetBalance())
}

type userBalance struct {
	sync.Mutex
	name    string
	balance int
}

func (u *userBalance) lock() {
	u.Mutex.Lock()
}
func (u *userBalance) unlock() {
	u.Mutex.Unlock()
}
func (u *userBalance) change(amount int) {
	u.balance = u.balance + amount
}
func tf(src *userBalance, dst *userBalance, amount int) {
	src.lock()
	fmt.Println("lock src ", src.name)
	src.change(-amount)

	time.Sleep(2 * time.Second)

	dst.lock()
	fmt.Println("lock dst", dst.name)
	dst.change(amount)

	time.Sleep(2 * time.Second)

	src.unlock()
	dst.unlock()
}

func TestDL(t *testing.T) {
	user1 := userBalance{
		name:    "andi",
		balance: 5000,
	}
	user2 := userBalance{
		name:    "dani",
		balance: 2000,
	}
	//andi=4000,dani=3000
	go tf(&user1, &user2, 1000)
	//andi=4500,dani=2500
	go tf(&user2, &user1, 500)
	time.Sleep(4 * time.Second)

	fmt.Println("balance ", user1.name, " = ", user1.balance)
	fmt.Println("balance ", user2.name, " = ", user2.balance)

	/*
			outputnya
		=== RUN   TestDL
		lock src  andi
		lock src  dani
		balance  andi  =  4000
		balance  dani  =  1500
		--- PASS: TestDL (4.01s)

		harusnya balance terakhirnya andi=4500,dani=2500
		ini terjadi karena di fungsi tf terjadi lock ke user1 dan user2 yang baru akan di unlock di akhir fungsi
		ketik groutine 1 lock user1, kemudian ingin lock user2.
		user2 sudah terburu di lcok oleh goroutine 2.
		harus berhati hati dalam penempatan lock dan unlock,apa lagi jika menggunakan defer.karena bisa menyebabkan
		deadlock
		di kode di atas, deadlock bisa di atasi dengan memindahkan unlock di akhir setiap change
	*/

}

func run(w *sync.WaitGroup, i int) {
	defer w.Done()
	fmt.Println("hello dari goroutine ke ", i)
}
func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		group.Add(1)
		go run(group, i)
	}
	group.Wait()
	fmt.Println("done")
}

func TestOnce(t *testing.T) {
	var once sync.Once
	var w sync.WaitGroup

	for i := 0; i < 10; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			once.Do(func() {
				fmt.Println("goroutine ke ", i)
			})
		}()

	}
	w.Wait()
	fmt.Println("done")
}

func TestPool(t *testing.T) {
	/*
		kita mendefinisikan fungsi untuk field new,field new ini bagian dari struct pool
		karena ketika pool.get di panggil dan isinya kosong. fungsi yang ada di new ini akan di panggil.
		cara set nya sama seperti mengeset field di dalam struct(jangan lupa untuk menambahkan koma di akhir)
	*/
	var pool = sync.Pool{
		// New : func() interface{} {
		// 	return "default job"
		// },
	}
	var w sync.WaitGroup
	text1 := "jobs 1"
	text2 := "jobs 2"
	text3 := "jobs3 "
	pool.Put(&text1)
	pool.Put(&text2)
	pool.Put(&text3)

	for i := 0; i < 10; i++ {
		w.Add(1)
		go func() {
			data := pool.Get()
			fmt.Println("goroutine ke", i, " dapat ", data)
			w.Done()
		}()
	}
	w.Wait()

}

func addData(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()
	data.Store(value, value)

}
func TestSyncMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		group.Add(1)
		go addData(data, i, group)
	}

	group.Wait()
	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})

	d, _ := data.Load(3)
	fmt.Println(d)
	data.Delete(3)
	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}

func WaitCondition(value int, cond *sync.Cond, group *sync.WaitGroup) {
	cond.L.Lock() //harus di lock sebelum menjalankan wait
	cond.Wait()   //akan menunggu sampai signal() atau broadcast() di jalankan
	fmt.Println("done ", value)
	cond.L.Unlock()
	group.Done()
}
func TestCond(t *testing.T) {
	locker := sync.Mutex{}
	cond := sync.NewCond(&locker)
	group := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		group.Add(1)
		go WaitCondition(i, cond, &group)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Second)
			cond.Signal()
		}
	}()
	group.Wait()

}

func TestAtomic(t *testing.T) {
	var data atomic.Value
	var wg sync.WaitGroup
	//inisialisasi
	text := "data awal"
	data.Store(text)
	//mengambil data
	d := data.Load().(string)
	fmt.Println("isi data = ", d)
	//menukar data (swap)
	new := "data baru"
	old := data.Swap(new).(string)
	fmt.Println("data lama = ", old)
	fmt.Println("data baru = ", data.Load().(string))
	//membandingkan dan menukar data (compare and swap)
	//bedanya jika data yang di bandingkan hasilnya tidak sama maka tidak akan di tukar
	expect := "data baru"
	new = "data paling baru"
	data.CompareAndSwap(expect, new)
	fmt.Println("data sesudah di CAS = ", data.Load().(string))

	var x int64
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("counter : ", x)
}

func TestTimer(t *testing.T) {
	timer := time.NewTimer(3 * time.Second)
	fmt.Println(time.Now())
	data := <-timer.C
	fmt.Println(data)
	//time.after
	channel := time.After(3 * time.Second)
	data = <-channel
	fmt.Println(data)
	//time.afterFunction
	var wg sync.WaitGroup
	wg.Add(1)
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("fungsi berhasil di jalankan")
		wg.Done()
	})

	wg.Wait()
}

func TestTicker(t *testing.T) {
	// Membuat ticker dengan interval 1 detik
	ticker := time.NewTicker(1 * time.Second)
	timer := time.After(5 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
		fmt.Println("thicker di hentikan")
	}()

loop:
	for {
		select {
		case <-ticker.C:
			fmt.Println("data masuk dari channel ticker")
		case <-timer:
			fmt.Println("data masuk dari channel timer")
			break loop
		}
	}
	fmt.Println("selesai")
}

func TestTickertick(t *testing.T) {
	tick := time.Tick(2 * time.Second)

	/*
		time.tick memang sederhana, namun ini akan terus menerus mengirim data tanpa bisa di stop.
		meskipung menggunakan tehnik label untuk keluar dari loop, time.tick masih akan terus berjalan
	*/
	for data := range tick {
		fmt.Println("data : ", data)
	}
}

func TestGomaxprocs(t *testing.T) {
	//menampilkan jumlah thread yang berjalan
	//jika memasukan angka di atas akang nol, maka akan di anggap melakukan perubahan
	//settingan GOMAXPROCS
	fmt.Println("jumlah thread :",runtime.GOMAXPROCS(0))	
	//mendapatkan jumlah cpu
	fmt.Println("jumlah core CPU :",runtime.NumCPU())
	//meskipun sedang tidak menjalankan code goroutine, secara default ada 2 goroutine yang berjalan
	//satu untuk garbage colection dan satu lagi untuk menjalankan program golang kita
	fmt.Println("jumlah Goroutine :",runtime.NumGoroutine())
}
