package main

import "fmt"

func numberBasics() {
	var n uint8 // uint8, uint16, uint32, uint, uint64
	n = 255

	fmt.Println("Value of n is", n)
	n += 100 // cycles with in the range (which is 0 to 255)
	fmt.Println("Value of n is", n)
	n -= 120
	fmt.Println("Value of n is", n)

	var n2 int16
	n2 = 30000
	fmt.Println("n2 is", n2)
	n2 += 10000 // no overflow; cycles with in the range (-32768 to 32767)
	fmt.Println("n2 is", n2)
}
