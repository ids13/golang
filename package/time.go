package main

import (
	"fmt"
	"time"
)

func main() {
	// Membuat Waktu:
	// time.Now(): Membuat waktu saat ini.
	// time.Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location): Membuat waktu dengan nilai yang ditentukan.
	now := time.Now()
	customTime := time.Date(2024, time.February, 14, 12, 0, 0, 0, time.UTC)
	fmt.Println("==MEMBUAT WAKTU==")
	fmt.Println("Waktu saat ini:", now)
	fmt.Println("Waktu kustom:", customTime)
	// Manipulasi Waktu:
	// t.Add(d Duration) Time: Menambahkan durasi ke waktu t.
	// t.Sub(u Time) Duration: Menghitung durasi antara dua waktu.
	// t.Format(layout string) string: Mengonversi waktu menjadi string dengan format tertentu.
	// t.Parse(layout string) timer.time,err : mengonversi string menjadi format waktu
	future := now.Add(time.Hour * 24)
	duration := future.Sub(now)
	fmt.Println("==MANIPULASI WAKTU==")
	fmt.Println("Waktu setelah satu hari:", future)
	fmt.Println("Durasi antara dua waktu:", duration)
	sekarang := time.Now()
	fmt.Println("sekaranng = ", sekarang.Format("2006-01-02"))
	tsekarang := "2002-12-02"
	psekarang, _ := time.Parse("2006-01-02", tsekarang) //karena mengghasilkan dua nilai jadi harus di tampung dulu
	fmt.Println(" 2002-12-02 di rubah ke format time = ", "tahun = ", psekarang.Year(), ", bulan = ", psekarang.Month(), ", tanggal = ", psekarang.Day())
	// Mengambil Komponen Waktu:
	// t.Year() int, t.Month() Month, t.Day() int: Mengambil tahun, bulan, dan hari dari waktu t.
	// t.Hour() int, t.Minute() int, t.Second() int: Mengambil jam, menit, dan detik dari waktu t.
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println("==KOMPONEN WAKTU==")
	fmt.Printf("Tahun: %d, Bulan: %s, Hari: %d, Jam: %d, Menit: %d, Detik: %d\n", year, month, day, hour, minute, second)
	// Pembandingan Waktu:
	// t.Before(u Time) bool, t.After(u Time) bool: Memeriksa apakah waktu t sebelum atau setelah waktu u.
	// t.Equal(u Time) bool: Memeriksa apakah waktu t sama dengan waktu u.
	t1 := time.Date(2024, time.February, 14, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, time.February, 14, 12, 0, 0, 0, time.UTC)
	before := t1.Before(t2)
	after := t1.After(t2)
	equal := t1.Equal(t2)
	fmt.Println("==PERBANDINGAN WAKTU==")
	fmt.Println("Sebelum:", before)
	fmt.Println("Setelah:", after)
	fmt.Println("Sama:", equal)
	// Zona Waktu:
	// time.LoadLocation(name string) (*Location, error): Memuat lokasi dari zona waktu dengan nama tertentu.
	// time.FixedZone(name string, offset int) *Location: Membuat zona waktu tetap dengan offset yang ditentukan.
	// FixedZone menghasilkan zona waktu sedangkan add menambah durasi waktu
	// Offset di hitung dalam detik,manipulasi ini untuk menghasilkan nilai yang di inginkan
	offsetSeconds := 3 * 60 * 60 // 10,800 detik setara 3 jam
	moscowZone := time.FixedZone("MSK", offsetSeconds)
	timeInMoscow := time.Now().In(moscowZone)
	loc, _ := time.LoadLocation("America/New_York")
	fmt.Println("==ZONA WAKTU==")
	fmt.Println("Waktu di Moscow:", timeInMoscow)
	fmt.Println("Lokasi waktu:", loc)
	// Konstanta Waktu:
	// time.Second, time.Minute, time.Hour: Konstanta untuk durasi satu detik, satu menit, dan satu jam.
	// time.RFC3339, time.RFC822: Konstanta format untuk RFC3339 (ISO 8601) dan RFC822 (RFC 822/1123) format.
	oneHour := time.Hour
	layout := time.RFC3339
	fmt.Println("==KONSTANTA WAKTU==")
	fmt.Println("Satu jam:", oneHour)
	fmt.Println("Layout RFC3339:", layout)
	// Durasi:
	// time.Duration: Representasi durasi waktu, digunakan untuk menambah atau mengurangkan waktu.
	durationExample := 2 * time.Hour
	fmt.Println("==DURASI==")
	fmt.Println("Durasi 2 jam:", durationExample)

	dur1 := time.Duration(2) * time.Hour
	dur2 := time.Duration(120) * time.Minute

	if dur1 == dur2 {
		fmt.Println("Durasi sama")
	}
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
}
