package main

import "fmt"

// Node adalah struktur untuk merepresentasikan node dalam linked list
type Node struct {
	data int
	next *Node
}

// prepend menambahkan elemen di depan linked list
func prepend(head **Node, value int) {
	/* kita bypass alamat dari variable pointer head.
	   karena pointer to pointer. meskipun bekerja dengan struct kita perlu menambahkan tanda *
	   untuk awalan kita membypass variable head *Node yang isi pointernya nil
	   kemudian kita membuat node baru yang isi data sesuai value yang di bypass, dan isi next nya adalah alamat
	   pointer dari yang di tampung di variable pointer head yaitu nil
	   lalu kita set isi variable pointer head dengan alamat node baru.
	   jadi setelah fungsi ini di jalankan isi variable node adalah alamat node baru

	*/
	newNode := &Node{data: value, next: *head}
	*head = newNode
}

// printList mencetak isi linked list
func printList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("NULL")
}

func main() {
	// Inisialisasi linked list kosong
	var head *Node

	// Menambahkan elemen ke linked list
	prepend(&head, 15)
	println(head)
	prepend(&head, 10)
	println(head)
	prepend(&head, 5)
	println(head)
	// Mencetak linked list
	fmt.Print("Linked List: ")
	printList(head)

	// Menunjukkan alamat dari pointer head
	fmt.Printf("Alamat dari pointer head: %p\n", &head)
}
