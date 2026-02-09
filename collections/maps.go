package main

import (
	"fmt"
	"sort"
)

func mapsGo() {
	states := make(map[string]string)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println("State of:", california)

	delete(states, "OR")
	fmt.Println(states)

	states["NY"] = "New York"
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%s: %s\n", k, v)
	}

	keys := make([]string, 0, len(states))
	for k := range states {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i := 0; i < len(keys); i++ {
		fmt.Printf("%s: %s\n", keys[i], states[keys[i]])
	}
}
