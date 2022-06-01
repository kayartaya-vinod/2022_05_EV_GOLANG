# Programming constructs (if-else, switch, for) and functions

### Exercise 1

Write a function to accept an int and check if it is a prime number or not. Return a boolean value representing the same.

### Exercise 2

Write a function that takes two int inputs (from and to) and returns the sum of all prime numbers between them.

### Exercise 3

Write a function that accepts a positive int input and returns all fibonacci elements until that number. For example, if the input is 7, then the return value must be a collection [0, 1, 1, 2, 3, 5, 8]

### Exercise 4

Write a function that accepts an int (0 < n < 10) and prints the following pattern (example shown is for 5):

```
*
**
***
****
*****
```

### Exercise 5

In trignometry, the Sine of an angle is represented by the series below:

![](https://camo.githubusercontent.com/c977e571a3af6b385289cf29629624d60f8c76ca195673fc575c364e087acb78/68747470733a2f2f77696b696d656469612e6f72672f6170692f726573745f76312f6d656469612f6d6174682f72656e6465722f7376672f33643936383930373763653339353239653331393832313339303132363165626366343831393764)

Write a function that accepts angle in degrees and returns the sine of the given angle. Call the function in main, multiple times by supplying multiple values and verify the same.

PS:

-   Divide the function into small reusable functions, if possible.
-   Do not use builtin Java classes like Math

### Exercise 6

Write a function that takes month and year as inputs, and prints the calendar for that month. If inputs are invalid, appropriate error message/s should be printed.

Sample output for the inputs (8, 2018):

```
Su Mo Tu We Th Fr Sa
          1  2  3  4
 5  6  7  8  9 10 11
12 13 14 15 16 17 18
19 20 21 22 23 24 25
26 27 28 29 30 31
```

### Exercise 7

Write a function called "sumOfEvensAndOdds", that takes a collection of integers as input and returns another collection of integers of length 2. The first element in the returned array is the sum of all even numbers in the input array, and the second element in the returned array is the sum of all odd numbers in the input array.

For example,

```go
var nums  = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
result := sumOfEvensAndOdds(nums);
// result should be equal to [30, 25]
```
