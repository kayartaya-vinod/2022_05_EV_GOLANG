package main

import (
	"example-13-unittesting/mathpack"
	"fmt"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from this panic", r)
		}
	}()

	s, err := mathpack.Sum("10, 20, asd")
	handleError(err)
	fmt.Println("Sum of 10, 20 is", s)
}
