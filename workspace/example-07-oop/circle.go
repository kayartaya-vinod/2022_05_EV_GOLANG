package main

import "math"

type circle struct {
	radius float64
}

func (this circle) area() float64 {
	return math.Pi * this.radius * this.radius
}

func (this circle) perimeter() float64 {
	return 2 * math.Pi * this.radius
}

func (this circle) info() string {
	return "type circle"
}
