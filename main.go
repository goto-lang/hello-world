package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")

	name := "Mik"
	msg, err := Hello(name)
	if err == nil {
		println(msg)
	}

	pointer := StringPointer(name)
	println(Type(pointer))
	println(Debug(pointer))
}

