package main

import "fmt"

/*
Using the following operators, write expressions and assign their values to variables:
g. ==
h. <=
i. >=
j. !=
k. <
l. >
Now print each of the variables
*/

var (
	equality         bool
	lessThanEqual    bool
	greaterThanEqual bool
	notEquality      bool
	lessThan         bool
	greaterThan      bool
)

func main() {
	fmt.Println("Exercise 2 - Ninja level 2")

	var num1 int = 20
	var num2 int = 25

	equality = num1 == num2
	lessThanEqual = num1 <= num2
	greaterThanEqual = num1 >= num2
	notEquality = num1 != num2
	lessThan = num1 < num2
	greaterThan = num1 > num2

	fmt.Printf("%d == %d -> %t\n%d <= %d -> %t\n%d >= %d -> %t\n%d != %d ->%t\n%d < %d ->%t\n%d > %d ->%t\n", num1, num2, equality, num1, num2, lessThanEqual, num1, num2, greaterThanEqual, num1, num2, notEquality, num1, num2, lessThan, num1, num2, greaterThan)

	// Above is not the efficient code as it has so much copy pasting

}
