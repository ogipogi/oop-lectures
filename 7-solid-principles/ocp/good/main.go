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

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Calculator struct{}

func (c Calculator) SumAreas(shapes []Shape) float64 {
	var total float64
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func main() {
	shapes := []Shape{
		Square{side: 2},
		Circle{radius: 3},
		Rectangle{width: 2, height: 3},
	}
	calculator := Calculator{}
	fmt.Println("Total area:", calculator.SumAreas(shapes))
}
