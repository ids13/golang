package hello

import "example.com/world"

func Hello() {
	println("hello", world.Message())
}
