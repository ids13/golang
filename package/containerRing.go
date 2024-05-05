package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

// ring ini bentuknya seperti cincin. jika di next sampai ujung.ketika di next lagi maka akan kembali lagi ke bagian awal
func main() {
	data := ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "value " + strconv.Itoa(i)
		data = data.Next()
	}
	data.Do(func(v interface{}) {
		fmt.Println(v)
	})
}
