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

type OutPrinter struct{}

func (op OutPrinter) PrintArea(shape Shape) {
	fmt.Printf("The area is: %f\n", shape.Area())
}

func (op OutPrinter) PrintAreaSum(sum float64) {
	fmt.Printf("The sum of the areas is: %f\n", sum)
}

type SumCalculator struct {
	shapes []Shape
}

func (sc SumCalculator) addShape(shape Shape) {
	sc.shapes = append(sc.shapes, shape)
}

func (sc SumCalculator) Sum() float64 {
	var sum float64
	for _, shape := range sc.shapes {
		sum += shape.Area()
	}
	return sum
}

func main() {
	square := Square{side: 2}
	circle := Circle{radius: 3}
	rectangle := Rectangle{width: 2, height: 3}

	printer := OutPrinter{}
	printer.PrintArea(square)
	printer.PrintArea(circle)
	printer.PrintArea(rectangle)

	sumCalculator := SumCalculator{}
	sumCalculator.addShape(square)
	sumCalculator.addShape(circle)
	sumCalculator.addShape(rectangle)

	sum := sumCalculator.Sum()
	printer.PrintAreaSum(sum)
}
