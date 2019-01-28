package main

import "fmt"

func builder() func(int) int {
	s := 0
	return func(x int) int {
		s += x
		return s
	}
}

func main() {
	sum := builder()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
	fmt.Printf("%T\n", sum)
}
