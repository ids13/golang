/*
	Buatlah program Golang yang dapat melakukan operasi dasar manajemen kontak,

seperti menambahkan kontak baru,
menampilkan daftar kontak,
dan mencari kontak berdasarkan nama.
*/
package main

import (
	"fmt"
)

// Contact adalah struktur data untuk merepresentasikan kontak
type Contact struct {
	Name  string
	Phone string
}

// ContactList adalah struktur data untuk menyimpan daftar kontak
type ContactList struct {
	Contacts []Contact
}

// AddContact menambahkan kontak baru ke daftar kontak
func (cl *ContactList) AddContact(name, phone string) {
	newContact := Contact{Name: name, Phone: phone}
	cl.Contacts = append(cl.Contacts, newContact)
}

// ShowContacts menampilkan daftar semua kontak
func (cl *ContactList) ShowContacts() {
	fmt.Println("Daftar Kontak:")
	for _, contact := range cl.Contacts {
		fmt.Printf("Nama: %s, Telepon: %s\n", contact.Name, contact.Phone)
	}
}

// SearchContact mencari kontak berdasarkan nama
func (cl *ContactList) SearchContact(name string) {
	found := false
	for _, contact := range cl.Contacts {
		if contact.Name == name {
			fmt.Printf("Nama: %s, Telepon: %s\n", contact.Name, contact.Phone)
			found = true
		}
	}

	if !found {
		fmt.Println("Kontak tidak ditemukan.")
	}
}

func main() {
	// Inisialisasi daftar kontak
	contacts := ContactList{}

	// Menambahkan beberapa kontak
	contacts.AddContact("John Doe", "123456789")
	contacts.AddContact("Jane Smith", "987654321")
	contacts.AddContact("Bob Johnson", "555555555")

	// Menampilkan daftar kontak
	contacts.ShowContacts()

	// Mencari kontak berdasarkan nama
	contacts.SearchContact("Jane Smith")
	contacts.SearchContact("Alice Johnson")
}
