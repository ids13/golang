package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

/*
	ada beberapa tipe data di golang yang memiliki implementasi io.writer dan io.reader

jadi bisa di jadikan paramter
contoh
byte.buffer
os.stdout /stdin / stderr
retunrn dari fungsi :
strings.NewReader
byte.NewBuffer
*/
func main() {

	buff := strings.NewReader("hallo dunia\naku siap")
	reader := bufio.NewReader(buff)

	for {
		//menggunakan readline untuk membaca bufio.reader
		x, _, err := reader.ReadLine()
		//pengecekan error
		if err != nil {
			//cek jika error nya EOF
			//method Error() akan mengembalikan nilai string, kita bandingkan dengan string EOF(end of file),
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("error :", err)
			break
		}
		println(string(x))
	}
	fmt.Println("===================")
	//mari gunakan bytes.buffer
	wadah := bytes.NewBuffer(nil)
	writer := bufio.NewWriter(wadah)
	_, _ = writer.WriteString("hallo dunia\n")
	_, _ = writer.WriteString("aku siap\n")
	_, _ = writer.WriteString("menulis\n")
	//karena bufio adalah io yang menggunakan buffered,ketika di write,tidak langsung di tulis ke buffer
	//jadi harus di flush
	writer.Flush()
	//sekarang gunakan scan untuk melihat isi buffer
	sc := bufio.NewScanner(wadah)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}

}
