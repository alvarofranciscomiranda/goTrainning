package main

import (
	"fmt"
	"gotrainning/collections"
)

func main() {
	colorNames := []string{"Red", "Green", "Blue"}
	hexValues := []int{0xFF0000, 0x00FF00, 0x0000FF}
	colors := collections.SlicesToObjects(colorNames, hexValues)
	fmt.Println(colors)
}
