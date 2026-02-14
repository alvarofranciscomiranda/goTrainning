package collections

import "fmt"

func StructFunc() {
	// Create an instance of the Dog struct
	myDog := Dog{
		Breed:  "Buddy",
		Weight: 5,
	}

	// Accessing struct fields
	println("My dog's Breed is:", myDog.Breed)
	println("My dog's Weight is:", myDog.Weight)

	// Modifying struct fields
	myDog.Weight = 6
	println("My dog's new Weight is:", myDog.Weight)

	fmt.Printf("My dog: %+v\n", myDog)
	fmt.Printf("My dog's Breed: %s, Weight: %d\n", myDog.Breed, myDog.Weight)
}

type Dog struct {
	Breed  string
	Weight int
}
