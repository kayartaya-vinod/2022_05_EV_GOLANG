package main

import "fmt"

func sliceBasics() {
	var nums = make([]int, 5)

	nums[0] = 1
	nums[1] = 2

	fmt.Println(nums)
	fmt.Printf("nums is %T\n", nums)

	nums = append(nums, 12, 34)
	fmt.Println("After appending, nums is", nums)

	nums2 := make([]int, len(nums))
	copy(nums2, nums)

	fmt.Println("copy of nums is", nums2)

	nums2 = append(nums, 99, 88, 44)
	fmt.Println("After appending to nums2, nums is", nums)
	fmt.Println("After appending to nums2, nums2 is", nums2)

}

func sliceOperations() {
	names := make([]string, 0)
	names = append(names, "Vinod", "Shyam", "Ramesh", "Harish", "Naveen", "Kiran", "Kishore")
	fmt.Println(names)
	fmt.Println(names[2])
	fmt.Println(names[2:5])
	fmt.Println(names[2:])
	fmt.Println(names[:5])

	fmt.Println("Vinod"[:3])

	var nums = [6]int{1, 2, 3, 4, 5, 5}
	fmt.Println(nums[:4])
}
