package reusable

import (
	"fmt"
	"time"
)

func GoroutineFunc() {
	go say("Hello from the Goroutine")
	println("Hello from Main")

	go func(message string) {
		fmt.Println(message)
	}("Hello from the anonymous Goroutine")

	time.Sleep(2 * time.Second)
	println("All done")
}

func say(message string) {
	time.Sleep(1 * time.Second)
	fmt.Println(message)
}
