package main

import "fmt"

type product struct {
	id    int
	name  string
	price float64
}

// methods - functions that are invoked using an object of a class (i.e, struct)

// the receiver can be a variable or a pointer
// if it is defined as a variable, then we cannot mutate the members
func (this product) print() {
	fmt.Printf("Product details: \n")
	fmt.Printf("ID     : %v\n", this.id)
	fmt.Printf("Name   : %v\n", this.name)
	fmt.Printf("Price  : %v\n", this.price)
	fmt.Println()
}

// since the receiver is a pointer, we will be able to mutate the field
// try by removing the asterisk
func (this *product) updatePrice(newPrice float64) {
	this.price = newPrice
}

func main2() {
	p1 := product{12, "Chai", 18.0}
	p1.print()
	p1.updatePrice(22.9)
	p1.print()
	p1.price = 29.1
	p1.print()
}
