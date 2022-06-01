package main

import "fmt"

func ifElseDemo() {

	var num int

	fmt.Print("Enter a number: ")
	_, err := fmt.Scanf("%d", &num)

	if err != nil {
		fmt.Println("There was an error:", err)
		return
	}

	if num%2 == 0 {
		fmt.Printf("%d is an even number\n", num)
	} else {
		fmt.Printf("%d is an odd number\n", num)
	}

}

func printMaxDays() {
	var month, year int
	fmt.Print("Enter month and year separated by a comma: ")
	fmt.Scanf("%d,%d", &month, &year)

	if month < 1 || month > 12 {
		fmt.Println("Month must be between 1 and 12, but received", month)
	} else if year < 1 {
		fmt.Println("Year must be more than zero, but received", year)
	} else {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			fmt.Println("The input month/year combo has 30 days")
		} else if month == 2 {
			if year%400 == 0 || (year%100 != 0 && year%4 == 0) {
				fmt.Println("The input month is february, year is a leap and has 29 days")
			} else {
				fmt.Println("The input month is february, year is a non-leap and has 28 days")
			}
		} else {
			fmt.Println("The input month/year combo has 31 days")
		}
	}
}
