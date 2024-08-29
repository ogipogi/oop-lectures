package main

type line interface {
	length() float64
}

type shape interface {
	line
	area() float64
}

type drawable interface {
	draw()
}

type object interface {
	shape
	volume() float64
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

type cube struct {
	square
}

func (c cube) volume() float64 {
	return c.side * c.side * c.side
}

func main() {
	s := square{side: 2}
	s.area()

	c := cube{square: square{side: 2}}
	c.area()
	c.volume()
}
