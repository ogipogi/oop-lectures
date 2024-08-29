package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type OutPrinter struct{}

func (op OutPrinter) PrintArea(shape Shape) {
	fmt.Printf("The area is: %f\n", shape.Area())
}

func main() {
	square := Square{side: 2}
	circle := Circle{radius: 3}
	printer := OutPrinter{}

	printer.PrintArea(square)
	printer.PrintArea(circle)
}
