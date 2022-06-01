package main

import "fmt"

func arrayBasics() {
	var nums [5]int
	fmt.Println(nums)

	nums[0] = 123
	nums[2] = 33
	// nums[12] = 40 // compilation error; invalid argument: array index 12 out of bounds
	fmt.Println(nums)

	fmt.Println("third element in the array is accessed using index 2")
	fmt.Printf("nums[%d] = %d\n", 2, nums[2])

	names := [6]string{"Vinod", "Shyam"}
	fmt.Println(names)

	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d] = \"%s\"\n", i, names[i])
	}

	fmt.Println("Using the range")
	for i, name := range names {
		fmt.Printf("i=%d, name=%s\n", i, name)
	}

	fmt.Println("names are (using for-each loop concept)...")
	for _, name := range names {
		fmt.Printf("'%s'\n", name)
	}

	fmt.Printf("names is %T\n", names)
}
