package reusable

import (
	"fmt"
	"strconv"
	"strings"
)

func ExerciseFunc() {
	calculation := calculate("  10.5 ", "5.2", "+")
	fmt.Printf("The result of the calculation is %v.\n", calculation)
}

func calculate(input1 string, input2 string, operation string) float64 {
	val1 := convertInputToValue(input1)
	val2 := convertInputToValue(input2)

	switch operation {
	case "+":
		return addValuesFloat(val1, val2)
	case "-":
		return subtractValues(val1, val2)
	case "*":
		return multiplyValues(val1, val2)
	case "/":
		return divideValues(val1, val2)
	}
	return 0
}

func convertInputToValue(input string) float64 {
	f1, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}
	return f1
}

func addValuesFloat(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	if value2 == 0 {
		panic("Cannot divide by zero")
	}
	return value1 / value2
}
