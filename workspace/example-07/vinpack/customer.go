package vinpack

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	HouseNo string `json:"houseNo"`
	Street  string `json:"street"`
	Area    string `json:"area"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type Customer struct {
	Id      int    `json:"id"`
	Name    string `json:"fullname"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address `json:"address"`
}

// This method represents a textual format of the Customer object.
// Automatically called when an object of this struct is treated like a string.
func (this Customer) String() string {
	return fmt.Sprintf("Customer(Id=%v, Name='%v', Email='%v', Phone='%v')",
		this.Id, this.Name, this.Email, this.Phone)
}

func (this Customer) Json() string {
	s, err := json.MarshalIndent(this, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(s)
}

func (this *Customer) FromJson(str string) {
	err := json.Unmarshal([]byte(str), &this)
	if err != nil {
		panic(err)
	}
}
