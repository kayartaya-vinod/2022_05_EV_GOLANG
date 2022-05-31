package mypack

import "fmt"

// note the lowercase "a" in the function name (camelCase)
// such methods are considered not exported (like private)
func aboutTrainer() {
	fmt.Println("Trainer name is Vinod Kumar Kayartaya")
}
