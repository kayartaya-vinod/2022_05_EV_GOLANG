package main

import "fmt"

func printFactorial() {
	var num int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scanf("%d", &num)

	if err != nil {
		fmt.Println("Error - ", err)
		return
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
