package main

import "fmt"

func usingConstants() {
	const pi float64 = 3.14157

	fmt.Println("Value of pi is", pi)

	radius := 12.34
	area := pi * radius * radius
	fmt.Println("Area of a circle with radius", radius, "is", area)

}
