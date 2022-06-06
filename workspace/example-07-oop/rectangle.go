package main

type rectangle struct {
	width   float64
	breadth float64
}

func (this rectangle) area() float64 {
	return this.width * this.breadth
}

func (this rectangle) perimeter() float64 {
	return 2 * (this.width + this.breadth)
}
