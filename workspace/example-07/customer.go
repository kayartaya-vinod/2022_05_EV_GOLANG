package main

import "fmt"

type address struct {
	houseNo string; 	street  string
	area    string; 	city    string
	state   string; 	country string
}

func (self address) printAddress() {
	fmt.Println("Address: ")
	fmt.Println(self.houseNo, self.street)
	fmt.Println(self.area, self.city)
	fmt.Println(self.state)
	fmt.Println(self.country)
}

type customer struct {
	id    int; 		name  string
	email string; 	phone string
	address // struct embedding; acts like inheritance
}

func (self customer) print() {
	fmt.Println("ID       :", self.id); fmt.Println("Name     :", self.name)
	fmt.Println("Email    :", self.email)
	fmt.Println("Phone    :", self.phone)
	self.address.printAddress()
}

func main3() {
	c1 := customer{id: 123, name: "Vinod", email: "vinod@vinod.co", phone: "9731424784"}
	c1.address = address{city: "Bangalore", state: "Karnataka", country: "India", area: "ISRO lyt"}
	c1.address.houseNo = "TF-1 Elegant Elite"
	c1.address.street = "1st cross, 1st main"
	c1.print()
	c1.printAddress() // via inheritance
}
