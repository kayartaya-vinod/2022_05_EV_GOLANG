package main

import "fmt"

// struct - typed collection of fields/variables.
// Useful for grouping data together - form records

type employee struct {
	id     int
	name   string
	salary int
}

func printEmployee(emp employee) {
	fmt.Printf("ID       : %v\n", emp.id)
	fmt.Printf("Name     : %v\n", emp.name)
	fmt.Printf("Salary   : %v\n", emp.salary)
	fmt.Println()
}

func main1() {
	e1 := employee{}
	e2 := employee{7788, "Scott", 45000}
	e3 := employee{salary: 34000, id: 2345, name: "Jones"}
	e4 := employee{name: "Ramesh", id: 8899}

	e5 := &e4
	e5.salary = 22200 // same as e4.salary=22200

	fmt.Printf("e1 = %v\n", e1)
	fmt.Printf("e2 = %v\n", e2)
	fmt.Printf("e3 = %v\n", e3)
	fmt.Printf("e4 = %v\n", e4)
	fmt.Printf("e5 = %v\n", e5)

	printEmployee(e1)
	printEmployee(e2)
	printEmployee(e3)
	printEmployee(e4)
	printEmployee(*e5)
}
