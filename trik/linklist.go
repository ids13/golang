package main

import "fmt"

// Node adalah struktur untuk merepresentasikan node dalam linked list
type Node struct {
	data int
	next *Node
}

// prepend menambahkan elemen di depan linked list
func prepend(head *Node, value int) *Node {
	newNode := &Node{data: value, next: head}
	return newNode
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
	head = prepend(head, 5)
	head = prepend(head, 10)
	head = prepend(head, 15)

	// Mencetak linked list
	fmt.Print("Linked List: ")
	printList(head)
}
