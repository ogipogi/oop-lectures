package main

type shape interface {
	area() float64
	volume() float64
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) volume() float64 {
	return 0
}

type cube struct {
	square
}

func (c cube) area() float64 {
	return c.side * c.side
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
