package main

import "fmt"

// variadic functions
// function that takes unknown number of paramters
func sum(nums ...int) (total int, err error) {

	if len(nums) < 2 {
		err = fmt.Errorf("Expected at least 2 inputs")
		return
	}
	total = 0
	for _, n := range nums {
		total += n
	}

	return
}

func subtotal(what int, nums ...int) float64 {
	switch what {
	case 1: // sum
		s, _ := sum(nums...)
		return float64(s)
	case 2: // average
		s, _ := sum(nums...)
		return float64(s) / float64(len(nums))
	default:
		return 0
	}
}

// recursion - a function calling itself
func factorial(num int) int {
	if num <= 1 {
		return 1
	}

	return num * factorial(num-1)
}

// f(5)
// return 5 * f(4)
//            return 4 * f(3)
//                       return 3 * f(2)
//                                  return 2 * f(1)
//                                             return 1
//                                  return 2 * 1
//                       return 3 * 2
//            return 4 * 6
// return 5 * 24
// 120

func main() {

	fmt.Println("Factorial of 5 is", factorial(5))
	fmt.Println("Factorial of 6 is", factorial(6))
	fmt.Println("Factorial of 15 is", factorial(15))

	s, err := sum(12, 34, 56, 67)
	if err == nil {
		fmt.Printf("s = %v\n", s)
	} else {
		fmt.Println("There was an error: ", err)
	}

	nums := []int{10, 20, 30, 40, 50, 68}
	s, _ = sum(nums...)
	fmt.Println("Sum of", nums, "is", s)

	sm := subtotal(1, nums...)
	fmt.Println("sum =", sm)

	avg := subtotal(2, nums...)
	fmt.Println("avg = ", avg)

	// anonynous functions
	// In Go, functions are first class objects
	// Can be passed as parameters to a function, and can be returned from a function (like int, float, string...)
	sayHello := func(name, city string) string {
		return fmt.Sprintf("Hello %s, how's weather in %s?", name, city)
	}

	sayHi := sayHello

	msg := sayHello("Vinod", "Bangalore")
	fmt.Println(msg)
	msg = sayHi("Shyam", "Shivamogga")
	fmt.Println(msg)

	fmt.Printf("Type of sayHello is %T\n", sayHello)

}
