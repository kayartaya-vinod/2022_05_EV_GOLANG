package main

import "fmt"

func readDateOfBirth() {
	var d int
	var m int
	var y int

	fmt.Print("Enter your date of birth in d/m/y format: ")
	fmt.Scanf("%d/%d/%d", &d, &m, &y)

	fmt.Printf("Please verify your date of birth (yyyy-m-d format) - %d-%d-%d\n\n", y, m, d)

}
