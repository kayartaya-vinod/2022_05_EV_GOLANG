package main

import (
	"fmt"
)

func add(num1 int, num2 int) int {
	return num1 + num2
}

func subtract(num1, num2 int) int {
	return num1 - num2
}

func divide(num1, num2 int) (int, int, string) {
	if num2 == 0 {
		return 0, 0, "Cannot divide by zero"
	}
	return num1 / num2, num1 % num2, ""
}

func factorial(num int) (int, error) {

	if num < 0 {
		return 0, fmt.Errorf("Cannot calculate factorial of %d", num)
	}

	f := 1
	for num > 1 {
		f *= num
		num--
	}
	return f, nil
}

func factorial2(num int) (fact int, err error) {
	if num < 0 {
		err = fmt.Errorf("Cannot calculate factorial of %d", num)
	} else {
		fact = 1
		for num > 1 {
			fact *= num
			num--
		}
	}
	return // current value of fact and err will be returned
}

func testFunctions() {

	f, e := factorial2(-5)
	fmt.Printf("f=%d, e=%v\n", f, e)

	// n := -12
	// f, err := factorial(n)

	// if err == nil {
	// 	fmt.Println("Factorial of", n, "is", f)
	// } else {
	// 	fmt.Println(err)
	// }

	// sum := add(12, 34)
	// diff := subtract(12, 34)
	// fmt.Println("sum is", sum)
	// fmt.Println("diff is", diff)

	// quotient, remainder, err := divide(123, 33)

	// if err == "" {
	// 	fmt.Println("123/44 results in quotient of", quotient, "and a remainder of", remainder)
	// } else {
	// 	fmt.Println("Error:", err)
	// }
}
