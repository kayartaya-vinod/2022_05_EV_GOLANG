package main

import (
	"fmt"
	"time"
)

//--------------------------------------------------------------------------------------------------
func printTable(num int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d X %d = %d\n", num, i, num*i)
		time.Sleep(200 * time.Millisecond)
	}
}

func main1() {
	fmt.Println("Start of main")
	go printTable(19)
	go printTable(876)
	fmt.Scanln()
	fmt.Println("End of main")
}

//--------------------------------------------------------------------------------------------------
func factorial(n int, ch chan int) {
	f := 1
	for i := 1; i <= n; i++ {
		f *= i
	}
	ch <- f // write to the channel
}

func main2() {
	ch := make(chan int)
	go factorial(10, ch) // call factorial in async style

	fact := <-ch // read from the channel
	fmt.Println("Factorial of 10 is", fact)
}

//--------------------------------------------------------------------------------------------------
func main3() {
	names := make(chan string)

	go func() {
		names <- "Vinod"
		names <- "Shyam"
	}()

	go func() {
		names <- "Kiran"
	}()

	// go func() {
	// 	time.Sleep(time.Second)
	// 	close(names)
	// }()

	// fmt.Println("names is", <-names)
	// fmt.Println("names is", <-names)
	// fmt.Println("names is", <-names)
	for n := range names {
		fmt.Println(n)
	}
}

//--------------------------------------------------------------------------------------------------
func prime(num int, ch chan int) {
	half := num / 2
	for i := 2; i <= half; i++ {
		if num%i == 0 {
			return
		}
	}
	ch <- num
}

func main4() {
	primes := make(chan int)
	for n := 1; n <= 500; n++ {
		go prime(n, primes)
	}

	fmt.Print("Primes from 1 to 500 are: ")
	for prime := range primes {
		fmt.Print(prime, ", ")
	}
	fmt.Println()
}

//--------------------------------------------------------------------------------------------------

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every half a second..."
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 seconds..."
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		// fmt.Println(<-c1)
		// fmt.Println(<-c2)
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)

		case msg2 := <-c2:
			fmt.Println(msg2)
		}

	}
}
