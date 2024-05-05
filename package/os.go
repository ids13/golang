package main

import (
	"fmt"
	"os"
)

/* os berguna untuk mengakses fitur sistem oprasi secara independen(berlaku semua oprating sistem) */

func main() {
	/* go run "c:\Users\langit\OneDrive\Documents\belajar\golang\src\package\os.go" naga langit
	jika arg nya tidak di tambahkan ketika di esekusi maka akan terjadi error routine*/
	arg := os.Args
	fmt.Println("arg ke 0 = ", arg[0])
	fmt.Println("arg ke 1 = ", arg[1])
	fmt.Println("arg ke 2 = ", arg[2])
	host, err := os.Hostname()
	fmt.Printf("host = %v error=%v\n", host, err)
	env := os.Getenv("USERNAME")
	fmt.Println("username = ", env)

}
