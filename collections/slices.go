package main

import (
	"fmt"
	"sort"
)

func slicesFunc() {
	// var colorsSlice = []string{"red", "green", "Blue"}

	var colors = make([]string, 0, 3)
	colors = append(colors, "red", "green", "blue")
	fmt.Println(colors)
	colors = append(colors, "purple", "fuschia")
	fmt.Println(colors)

	colors = remove(colors, 1)
	fmt.Println(colors)

	sort.Strings(colors)
	fmt.Println(colors)
}

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}
