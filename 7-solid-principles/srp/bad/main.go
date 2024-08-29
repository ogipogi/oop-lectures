package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (s Square) Area() {
	area := s.side * s.side
	fmt.Printf("The area of the square is: %f\n", area)
}

func (c Circle) Area() {
	area := math.Pi * c.radius * c.radius
	fmt.Printf("The area of the circle is: %f\n", area)
}

func main() {
	square := Square{side: 2}
	circle := Circle{radius: 3}

	square.Area()
	circle.Area()
}
