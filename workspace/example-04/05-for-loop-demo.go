package main

import (
	"fmt"
	"math"
)

func infiniteLoopDemo() {
	for {
		var num float64
		fmt.Print("Enter a number: ")
		fmt.Scanf("%f", &num)
		fmt.Printf("Square root of %f is %f\n", num, math.Sqrt(num))

		fmt.Print("Do you wish to try another? (yes/no) ")
		var ans string
		fmt.Scanf("%s", &ans)

		if ans == "no" {
			break
		}
	}
	fmt.Println("That's all folks!")
}

func printFactorial() {
	var num int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scanf("%d", &num)

	if err != nil {
		panic("There was an error: " + err.Error())
	}

	if num < 0 {
		panic("Not calculating factorial of negative input")
	}

	f := 1

	for num > 1 {
		f *= num
		num--
	}

	fmt.Println("Factorial for the input is", f)
}

func printMultiplicationTable() {
	var num int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scanf("%d", &num)

	if err != nil {
		fmt.Println("Error - ", err)
		return
	}

	for i := 1; i <= 15; i++ {
		fmt.Printf("%d X %d = %d\n", num, i, num*i)
	}

	fmt.Println("That's all folks!")
}
