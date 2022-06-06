package main

import (
	"fmt"
	"os"
	"strconv"
)

func feetToMeters() {
	arg := os.Args[1]
	feet, _ := strconv.ParseFloat(arg, 64)

	fmt.Println("feet is", feet)
	meters := feet * 0.3048
	fmt.Println("The same in meters is", meters)

}
