package main

import (
	"errors"
	"fmt"
)

var (
	valerr   = errors.New("validation error")
	notfount = errors.New("not found")
)

func name(name string) error {
	if name == "" {
		return valerr
	} else if name != "naga" {
		return notfount
	}
	return nil
}
func main() {
	status := name("naga")
	if errors.Is(status, valerr) {
		fmt.Println("perintah yang di jalankan jika validasi error")
	} else if errors.Is(status, notfount) {
		fmt.Println("perintah yang di jalankan jika user notfound")
	} else {
		fmt.Println("user sudah pas")
	}

}
