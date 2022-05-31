package mypack

import "fmt"

// note the capital "W" in the function name (PascalCase)
// such methods are considered to be exported (like public)
func Welcome() {
	fmt.Println("Welcome to Golang training")
	aboutTrainer()
	// aboutTrainer() is accessible to any other functions of the
	// same package
}
