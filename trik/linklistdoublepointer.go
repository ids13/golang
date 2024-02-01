package main

import "fmt"

// Node adalah struktur untuk merepresentasikan node dalam linked list
type Node struct {
	data int
	next *Node
}

// prepend menambahkan elemen di depan linked list
func prepend(head **Node, value int) {
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
	prepend(&head, 10)
	prepend(&head, 5)

	// Mencetak linked list
	fmt.Print("Linked List: ")
	printList(head)

	// Menunjukkan alamat dari pointer head
	fmt.Printf("Alamat dari pointer head: %p\n", &head)
}
