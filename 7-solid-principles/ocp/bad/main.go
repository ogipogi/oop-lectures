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

type Calculator struct{}

func (c Calculator) SumAreas(shapes []interface{}) float64 {
	var total float64
	for _, shape := range shapes {
		switch s := shape.(type) {
		case Square:
			total += s.side * s.side
		case Circle:
			total += math.Pi * s.radius * s.radius
		}
	}
	return total
}

func main() {
	shapes := []interface{}{
		Square{side: 2},
		Circle{radius: 3},
	}
	calculator := Calculator{}
	fmt.Println("Total area:", calculator.SumAreas(shapes))
}
