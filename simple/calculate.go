package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Uncomment this import section to use required Go packages

// calculate() returns the the result of adding 2 numbers
// after parsing them from strings
func calculate(value1 string, value2 string) float64 {

	f1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}
	f2, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}

	fmt.Printf("The sum is: %.2f\n", f1+f2)
	return f1 + f2
}
