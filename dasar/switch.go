package main

import "fmt"

func main() {
	/*
	   jika pengkondisian sederhana sebaiknya menggunakan switch case
	*/
	// penggunaan sederhana switch
	huruf := "x"
	switch huruf {
	case "a":
		fmt.Println("hurufnya a")
	case "b":
		fmt.Println("hurufnya b")
	case "c":
		fmt.Println("hurufnya c")
	default:
		fmt.Println("saya tidak tau hurufnya")
	}
	/*
		switch bisa di gunakan seperti if else (melakukan perbandingan)
		untuk melakukannya ,variable yang akan di bandingkan jangan di tulis setelah kata switch
	*/
	nilai := 65
	switch {
	case nilai > 70:
		fmt.Println("nilai di atas tujuh puluh")
	case nilai > 60:
		fmt.Println("nilai di atas enam puluh")
	case nilai >= 50:
		fmt.Println("nilai di atas lima puluh")
	case nilai < 50:
		fmt.Println("nilai di bawah lima puluh")
	default:
		fmt.Println("tolong masukan nilai")
	}

	/*
		fallthrough di gunakan untuk membuat proses pengecekan di lanjutkan ke case berikutnya tanpa menghiraukan kondisinya terpenuhi atau tidak
		karena pada dasarnya switch akan berhenti ketika case sudah terpenuhi.oleh karena itu
		gunakan fallthrough untuk memaksa program mengeksekusi case berukutnya*/

	ujian := 70
	switch {
	case ujian > 80:
		fmt.Println("mantap")
	case (ujian >= 50) && (ujian < 80):
		fmt.Println("lumayan")
		fallthrough
	case ujian <= 50: //padahal kondisi nya tidak terpenuhi tapi karena adanya fallthrouh .jadi di eksekusi
		fmt.Println("belajar lagi")
	default:
		fmt.Println("error")
	}
}
