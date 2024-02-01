package main

import "fmt"

func main() {
	var name1 [3]string
	name1[0] = "text1"
	name1[1] = "text2"
	name1[2] = "text3"

	var name2 = [3]int{
		1,
		2,
		3,
	}

	var name3 = [...]int{
		4,
		5,
		6,
	}

	var name4 = [2][3]int{{1, 2, 3}, {1, 2, 3}}

	fmt.Println(name1)
	fmt.Println(len(name1))
	fmt.Println(name2)
	fmt.Println(len(name2))
	fmt.Println(name3)
	fmt.Println(len(name3))
	fmt.Println(name4)
	fmt.Println(len(name4))
	fmt.Println(len(name4[0]))
}
