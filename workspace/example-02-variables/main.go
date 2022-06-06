package main

import (
	"fmt"
	"os"
)

func main() {

	var n1, n2, n3 int
	n1, n2, n3 = 10, 20, 30
	fmt.Println(n1, n2, n3)

	fmt.Println("len(os.Args) is", len(os.Args))
	fmt.Println(os.Args)
	fmt.Println(os.Args[1])

	var myName = "Vinod Kumar"
	var myCity = "Bangalore"

	fmt.Println("My name is", myName)
	fmt.Println("I live in", myCity)

	fmt.Printf("My name is %s and I live in %s\n", myName, myCity)
	fmt.Printf("My name is %v and I live in %v\n", myName, myCity)

	var myAge = 48 // type inference
	fmt.Printf("My age is %v years\n", myAge)
	fmt.Printf("myName is %T, myCity is %T, and myAge is %T\n", myName, myCity, myAge)

	// myAge = "asdf" // error; since the type of myAge was decided as int while declaring

	var myState string
	myState = "Karnataka"

	fmt.Printf("I belong to %v state", myState)

	myCountry := "India" // declaration + type inference
	// myCountry = "Bharath"
	fmt.Printf(" of %v\n", myCountry)

}
