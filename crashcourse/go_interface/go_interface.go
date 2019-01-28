package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	c1 := Circle{10}
	fmt.Println(getArea(c1))
}
