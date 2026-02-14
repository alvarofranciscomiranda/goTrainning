package collections

// slicesToObjects() returns a slice of Color objects
func SlicesToObjects(colorNames []string, hexValues []int) []Color {
	colors := make([]Color, 0, len(colorNames))
	for i := 0; i < len(colorNames); i++ {
		colors = append(colors, Color{Name: colorNames[i], Hex: hexValues[i]})
	}
	return colors
}

type Color struct {
	Name string
	Hex  int
}

/*
	colorNames := []string{"Red", "Green", "Blue"}
	hexValues := []int{0xFF0000, 0x00FF00, 0x0000FF}
	colors := collections.SlicesToObjects(colorNames, hexValues)
	fmt.Println(colors)
*/
