package main

import "fmt"

type Person struct {
	ID   int
	Name string
	Age  int
}

func main() {
	// Slice untuk menyimpan data
	var people []Person

	// Map untuk mengindeks data berdasarkan ID
	peopleByID := make(map[int]*Person)

	// Menambahkan data ke "database"
	person1 := Person{ID: 1, Name: "John", Age: 25}
	people = append(people, person1)
	peopleByID[person1.ID] = &person1

	person2 := Person{ID: 2, Name: "Jane", Age: 30}
	people = append(people, person2)
	peopleByID[person2.ID] = &person2

	// Mengakses data menggunakan map
	idToRetrieve := 1
	retrievedPerson, exists := peopleByID[idToRetrieve]
	if exists {
		fmt.Printf("Retrieved Person with ID %d: %v\n", idToRetrieve, retrievedPerson)
	} else {
		fmt.Printf("Person with ID %d not found\n", idToRetrieve)
	}

	// Menampilkan data dari "database"
	for _, person := range people {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", person.ID, person.Name, person.Age)
	}
}
