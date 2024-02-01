package main

import "fmt"

func add(t *string) {
	*t = "hallo " + *t
}
func main() {
	nama := "naga"
	add(&nama)
	fmt.Println(nama)
}
