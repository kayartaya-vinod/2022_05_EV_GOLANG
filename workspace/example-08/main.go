package main

import (
	"encoding/json"
	"fmt"

	"github.com/kayartaya-vinod/golang-tags-demo/vinpack"
)

func main1() {
	c1 := vinpack.Customer{Id: 1, Name: "Vinod", Email: "vinod@vinod.co", Phone: "9731424784"}
	c1.Address = vinpack.Address{City: "Bangalore", State: "Karnataka", Country: "India"}
	fmt.Println("c1 is", c1)

	fmt.Println(c1.Json())
	jsonStr := `{"id":7788,"fullname":"John Miller","email":"john.miller@xmpl.com","phone":"5553452987"}`

	c2 := vinpack.Customer{}
	fmt.Println("(before unmarshall) c2 is", c2)
	c2.FromJson(jsonStr)
	fmt.Println("(after unmarshall) c2 is", c2)

	jsonStr = `{
		"name": "Vinod",
		"age": 48,
		"isMarried": true,
		"phones": ["9731424784", "9844083934"],
		"address": {
			"city": "Bangalore",
			"state": "Karnataka"
		}
	}`

	// interface{} represents any datatype
	var p1 map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &p1)
	if err != nil {
		panic(err)
	}
	fmt.Println(p1)
	fmt.Println("Name is", p1["name"])
	fmt.Println("Age is", p1["age"])
	fmt.Println("isMarried =", p1["isMarried"])
	fmt.Println("Phones is", p1["phones"])
	fmt.Println("Address is", p1["address"])

}
