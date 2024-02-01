package main

import "fmt"

/* dengan interface anda bisa memasukan struct mahluk yang berbeda beda.untuk fungsi yang sama
(tidak perlu memodifikasi fungsinya) */
// Buat interface
type Speaker interface {
	Speak() string
}

// Implementasikan interface untuk manusia
type Human struct {
	Name string
}

// membuat kontrak untuk ke interface Speaker,untuk fungsi Speak() yang ada di dalamnya
func (h Human) Speak() string { //format return dan parameternya harus sama
	return "Hello, my name is " + h.Name
}

// Implementasikan interface untuk kucing
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow, I'm " + c.Name
}

// Fungsi untuk mendengarkan pembicara apa pun
func Listen(speaker Speaker) {
	fmt.Println(speaker.Speak())
}

func main() {
	// Buat objek manusia
	person := Human{Name: "John"}

	// Buat objek kucing
	cat := Cat{Name: "Whiskers"}

	// Dengarkan manusia dan kucing
	Listen(person)
	Listen(cat)
}
