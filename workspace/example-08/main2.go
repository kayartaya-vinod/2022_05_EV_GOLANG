package main

import (
	"encoding/xml"
	"fmt"

	"github.com/kayartaya-vinod/golang-tags-demo/vinpack"
)

func main() {
	e1 := vinpack.Employee{Id: 1122, Name: "James Scott", Salary: 34500, Department: "ADMIN"}

	e1Xml, _ := xml.MarshalIndent(e1, "", "   ")
	fmt.Println(xml.Header + string(e1Xml))
}
