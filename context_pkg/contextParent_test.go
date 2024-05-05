package contextpkg

import (
	"context"
	"fmt"
	"testing"
)

func TestValue(t *testing.T) {
	/*
	   A
	   ├── B
	   │   ├── D
	   │   └── E
	   └── C
	       ├── F
	       └── G

	*/
	contextA := context.Background()
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")
	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")
	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextC, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	// jika tidak di temukan, maka akan mencari ke parent nya
	fmt.Println(contextF.Value("f")) //langsung dapat
	fmt.Println(contextF.Value("c")) //di dapat dari parent
	fmt.Println(contextF.Value("b")) //tidak dapat, karena beda parent
	fmt.Println(contextA.Value("b")) //tidak dapat, karena tidak bisa mengambil dari child
}
