package main

import "fmt"

func fn1(num int) { // num is a new variable, each time this function is called
	num += 10
	fmt.Println("After adding 10, num is", num)
}

func fn2(num *int) { // num is a new variable, but of a pointer type. It holds the addres of the value being passed
	// fmt.Println("num is", num)
	*num += 10 // change the value in the address held by num
	fmt.Println("After adding 10, *num is", *num)
}

func sequence() func() int {
	x := 1
	fmt.Println("x is", x)
	num := 0
	return func() int {
		num++
		return num
	}
}

func calculator() (func(int) int, func(int) int) {
	var num int

	return func(n int) int {
			num += n
			return num
		},
		func(n int) int {
			num -= n
			return num
		}
}

func main() {

	add, subtract := calculator()
	// the closure scope which contains "num" is shared across all functions returned by calcualtor()
	fmt.Println(add(5))
	fmt.Println(add(50))
	fmt.Println(add(45))
	fmt.Println(subtract(10))
	fmt.Println(subtract(20))

	nextNum := sequence()
	fmt.Printf("nextNum is %T\n", nextNum)
	// Once execution of the function sequence() is over, all the local variables declared inside that
	// should be removed. However, in the nextNum() function, "num" is used, but not declated. So, the variable "x"
	// in the sequence() function, is removed from the stack, but the variable "num" in the same function is retained
	// in a special scope called "closure"
	n := nextNum() // increments the value in the closure scope and returns the same
	fmt.Println("n =", n)
	n = nextNum()
	fmt.Println("n =", n)
	n = nextNum()
	fmt.Println("n =", n)
	n = nextNum()
	fmt.Println("n =", n)

	// fmt.Println("num =", num) // num is not a gloabl variable; cannot be accessed anywhere else other than nextNum()

	// n := 15
	// fn1(n) // pass by value (copy)
	// fmt.Println("In main, n =", n)
	// fn1(n)
	// fmt.Println("In main, n =", n)
	// fn1(n)
	// fmt.Println("In main, n =", n)

	// m := 15
	// // fmt.Println("Address of n is", &m)
	// fn2(&m) // pass by reference/address
	// fmt.Println("In main, m =", m)
	// fn2(&m)
	// fmt.Println("In main, m =", m)
	// fn2(&m)
	// fmt.Println("In main, m =", m)
}
