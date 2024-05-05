package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("percobaan.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode())
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	defer file.Close()

	i, err := file.WriteString("Ini percobaan membuat file Golang")
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	fmt.Println("i =", i)
}
