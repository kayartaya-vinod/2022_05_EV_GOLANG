package main

import "fmt"

func mapsBasics() {
	emps := make(map[int]string) // key is int (empno) and value is string (name)
	emps[7788] = "Scott"
	emps[7654] = "James"
	emps[7896] = "Allen" // adds as a new value
	emps[7788] = "James" // replaces the existing value

	fmt.Println(emps)

	emp1 := emps[7788]
	fmt.Println(emp1)

	id := 7896
	emp2, found := emps[id]
	if found {
		fmt.Println(emp2)
	} else {
		fmt.Println("No employee found for id", id)
	}

	studentMarks := map[string]int{"Kiran": 98, "Kishore": 78, "Harish": 75}
	for name, marks := range studentMarks {
		fmt.Printf("%s has got %d marks\n", name, marks)
	}
	fmt.Println("Student names...")
	for name := range studentMarks {
		fmt.Println(name, "-->", studentMarks[name])
	}

	fmt.Println("Student marks...")
	for _, marks := range studentMarks {
		fmt.Println(marks)
	}
}
