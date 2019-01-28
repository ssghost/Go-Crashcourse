package main

import "fmt"

func main() {
	s := "Hello"
	p := &s

	fmt.Printf("%T\n", &s)
	fmt.Printf("%T\n", *p)
}
