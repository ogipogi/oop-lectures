package main

import "fmt"

type transport interface {
	getName() string
	drive()
	fly()
}

// base entity type
type vehicle struct {
	name string
}

func (v vehicle) getName() string {
	return v.name
}

// sub entity type
type car struct {
	vehicle
	wheel int
	gate  int
}

func (c car) drive() {
	fmt.Println("Driving car...")
}

func (c car) fly() {
	// nothing
}

type motorcycle struct {
	vehicle
	wheel int
}

func (m motorcycle) drive() {
	fmt.Println("Driving motorcycle...")
}

func (m motorcycle) fly() {
	// nothing
}

type printer struct{}

func (p printer) printTransportName(t transport) {
	fmt.Println("Transport name: ", t.getName())
}

func main() {
	car := car{
		vehicle: vehicle{name: "Car"},
		wheel:   4,
		gate:    4,
	}

	motorcycle := motorcycle{
		vehicle: vehicle{name: "Motorcycle"},
		wheel:   2,
	}

	printer := printer{}
	printer.printTransportName(car)
	printer.printTransportName(motorcycle)
}
