/* ini adalah main dari modul dengan nama paket */
package main

import (
	"fmt"
	"paket/say"
)

func main() {
	say.HelperFunction()
	say.UtilFunction()
	a := say.Data{Name: "tes", Value: 13} //test
	fmt.Println(a.Name, a.Value)
}
