package main

import (
	"fmt"
)

// an interface is a group of related function declarations
// any object that has implemented those functions, can be treated as
// objects of that interface

type measureable interface {
	area() float64
	perimeter() float64
}

// polymorphic function; behavior changes as per the object passed
func measure(m measureable) {
	fmt.Printf("Measuring the object of type %T\n", m)
	fmt.Println("Area is", m.area())
	fmt.Println("Perimeter is", m.perimeter())
	// cannot invoke object specific methods
	// can only invoke methods defined in the interface "measureable"
	fmt.Println()
}

func main() {
	c1 := circle{12.34}
	r1 := rectangle{12.34, 56.78}

	measure(c1)
	measure(r1)

}
