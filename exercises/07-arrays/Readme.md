# EXERCISE: Declare empty arrays

1. Declare and print the following arrays with their types:
1. The names of your three best friends (names array)
1. The distances to five different locations (distances array)
1. A data buffer with five bytes of capacity (data array)
1. Currency exchange ratios only for a single currency (ratios array)
1. Up/Down status of four different web servers (alives array)
1. A byte array that doesn't occupy memory space (zero array)
1. Print only the types of the same arrays.
1. Print only the elements of the same arrays.

### HINT

When printing the elements of an array, you can use the usual Printf verbs.

For example:

When printing a string array, you can use "%q" verb as usual.

### EXPECTED OUTPUT

```
names : [3]string{"", "", ""}
distances: [5]int{0, 0, 0, 0, 0}
data : [5]uint8{0x0, 0x0, 0x0, 0x0, 0x0}
ratios : [1]float64{0}
alives : [4]bool{false, false, false, false}
zero : [0]uint8{}

names : [3]string
distances: [5]int
data : [5]uint8
ratios : [1]float64
alives : [4]bool
zero : [0]uint8

names : ["" "" ""]
distances: [0 0 0 0 0]
data : [0 0 0 0 0]
ratios : [0.00]
alives : [false false false false]
zero : []
```

# EXERCISE: Get and Set Array Elements

1. Use the 01-declare-empty exercise
2. Remove everything but the array declarations
3. Assign your best friends' names to the names array
4. Assign distances to the closest cities to you to the distance array
5. Assign arbitrary bytes to the data array
6. Assign a value to the ratios array
7. Assign true/false values to the alives arrays
8. Try to assign to the zero array and observe the error
9. Now use ordinary loop statements for each array and print them (do not use for range)
10. Now use for range loop statements for each array and print them
11. Try assigning different types of values to the arrays, break things, and observe the errors
12. Remove some of the array assignments and observe the loop outputs (zero values)

### EXPECTED OUTPUT

Note: The output can change depending on the values that you've assigned to them, of course.
You're free to assign any values.

```
# names

names[0]: "Einstein"
names[1]: "Tesla"
names[2]: "Shepard"

# distances

distances[0]: 50
distances[1]: 40
distances[2]: 75
distances[3]: 30
distances[4]: 125

# data

data[0]: 72
data[1]: 69
data[2]: 76
data[3]: 76
data[4]: 79

# ratios

ratios[0]: 3.14

# alives

alives[0]: true
alives[1]: false
alives[2]: true
alives[3]: false

# zero

```

FOR RANGES

```

# names

names[0]: "Einstein"
names[1]: "Tesla"
names[2]: "Shepard"

# distances

distances[0]: 50
distances[1]: 40
distances[2]: 75
distances[3]: 30
distances[4]: 125

# data

data[0]: 72
data[1]: 69
data[2]: 76
data[3]: 76
data[4]: 79

# ratios

ratios[0]: 3.14

# alives

alives[0]: true
alives[1]: false
alives[2]: true
alives[3]: false

# zero
```

# EXERCISE: Currency Converter

In this exercise, you're going to display currency exchange ratios against USD.

1. Declare a few constants with iota. They're going to be the keys of the array.
2. Create an array that contains the conversion ratios. You should use keyed elements and the contants you've declared before.
3. Get the USD amount to be converted from the command line.
4. Handle the error cases for missing or invalid input.
5. Print the exchange ratios.

### EXPECTED OUTPUT

```
go run main.go
Please provide the amount to be converted.
```

```
go run main.go invalid
Invalid amount. It should be a number.
```

```
go run main.go 10.5
10.50 USD is 9.24 EUR
10.50 USD is 8.19 GBP
10.50 USD is 1186.71 JPY
```

```
go run main.go 1
1.00 USD is 0.88 EUR
1.00 USD is 0.78 GBP
1.00 USD is 113.02 JPY
```

# EXERCISE: Number Sorter

Your goal is to sort the given numbers from the command-line.

1. Get the numbers from the command-line.
2. Create an array and assign the given numbers to that array.
3. Sort the given numbers and print them.

### RESTRICTION

-   Maximum 5 numbers can be provided
-   If one of the arguments is not a valid number, skip it

### HINTS

-   You can use the bubble-sort algorithm to sort the numbers.
    Please watch this: https:youtu.be/nmhjrI-aW5o?t=7
-   When swapping the elements, do not check for the last element.
    Or, you will receive this error:
    "panic: runtime error: index out of range"

### EXPECTED OUTPUT

```
go run main.go
Please give me up to 5 numbers.
```

```
go run main.go 6 5 4 3 2 1
Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...
```

```
go run main.go 5 4 3 2 1
[1 2 3 4 5]
```

```
go run main.go 5 4 a c 1
[0 0 1 4 5]
```