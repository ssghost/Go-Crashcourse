package main

import "fmt"

type Person struct {
	name  string
	count int
}

func (p *Person) upvote() {
	p.count++
}

func main() {
	p1 := Person{"p1", 0}
	p1.upvote()
	fmt.Println(p1)
}
