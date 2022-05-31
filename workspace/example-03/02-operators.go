package main

import (
	"fmt"
	"strconv"
)

func operatorsDemo() {
	fmt.Println("Using arithmetic operators with number types")

	// when an int and float are used in an expression, the result
	// of the expression will be the one with the max capacity.
	// in this case, it will be float64
	fmt.Printf("%v %T\n", 8*-4.0, 8*-4.0)

	fmt.Printf("%v %T\n", -4/2, -4/2)

	// remainder operator
	fmt.Printf("%v %T\n", 123%12, 123%12)

	// changing sign of a value/variable
	var n1 int = -12
	fmt.Println(-n1, - -12, -(-12))

	var myAge, yourAge = 48, 23
	var averageAge float64

	averageAge = float64(myAge+yourAge) / 2
	fmt.Printf("myAge=%v, yourAge=%v, averageAge=%v\n", myAge, yourAge, averageAge)
	fmt.Printf("myAge=%T, yourAge=%T, averageAge=%T\n", myAge, yourAge, averageAge)

	var n int
	fmt.Println("n is", n) // n not assigned; defaults to 0

	n = n + 1
	n += 1 // applies to + - * / %
	n++

	fmt.Println("n is", n)
	fmt.Println(10 + 20)
	fmt.Println("Vinod " + "Kumar")
	fmt.Println("n is " + strconv.Itoa(n))
}
