package main

import "fmt"

/* interface adalah tipe data abstrak,jadi tidak bisa impelentasi secara langsung(tidak bisa berisi tipe data)
interface berisi tipe data
biasanya di gunakan sebagai kontrak*/

type hasname interface {
	getname() string
}

type data struct {
	nama string
}

func (n data) getname() string {
	return n.nama
}

/* parameter say kontraknya adalah hasname
jadi siapapun itu / tipe data apapun itu yang mempunyai function yang namanya */
func say(n hasname) {
	fmt.Println("hallo ", n.getname())
}

func main() {
	d := data{nama: "joko"}
	say(d)
}
