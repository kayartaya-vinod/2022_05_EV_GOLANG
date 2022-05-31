package main

import "fmt"

func usingStrings() {
	// string literal
	var s1 string
	s1 = "My name is Vinod.\nI live in Bangalore.\n\n"
	fmt.Println(s1)

	s2 := `
package mypack

import "fmt"

// note the lowercase "a" in the function name (camelCase)
// such methods are considered not exported (like private)
func aboutTrainer() {
	fmt.Println("Trainer name is Vinod Kumar Kayartaya")
}
`
	fmt.Println("----------------------------------------")
	fmt.Println(s2)
	fmt.Println("----------------------------------------")
}
