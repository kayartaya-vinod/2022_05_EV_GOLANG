package main

import "fmt"

func scanfDemo() {
	var name string
	var age int
	var height float64

	fmt.Print("Enter your name, age and height: ")
	fmt.Scanf("%s %d %f", &name, &age, &height)

	fmt.Printf("%s is %d years old and %.2f ft. tall\n", name, age, height)

}
