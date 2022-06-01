package main

import (
	"fmt"
	"time"
)

func greetByTime() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 16:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}

func checkForWeekday() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Todays is a weekend :-)")
	default:
		fmt.Println("Today is a weekday :-(")
	}
}

func validateDate() {
	var d, m, y int
	fmt.Print("Enter date in d/m/y format: ")
	_, err := fmt.Scanf("%d/%d/%d", &d, &m, &y)

	if err != nil {
		fmt.Println("There was an error:", err)
		return
	}
	if y <= 0 {
		fmt.Println("Invalid value for year; must be > 0")
		return
	}
	if m < 1 || m > 12 {
		fmt.Println("Invalid value for month; must be between 1 and 12")
		return
	}

	maxDays := 31

	switch m {
	case 4, 6, 9, 11:
		maxDays = 30
	case 2:
		if y%400 == 0 || (y%100 != 0 && y%4 == 0) {
			maxDays = 29
		} else {
			maxDays = 28
		}
	}

	if d < 1 || d > maxDays {
		fmt.Println("Invalid value for date; must be between 1 and", maxDays)
		return
	}

	fmt.Println("Input values for date, month and year forms a valid date!")

}
