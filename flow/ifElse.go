package flow

func IfFunc(theAnswer int) {
	if theAnswer < 50 {
		println("The answer is less than 50")
	} else if theAnswer == 42 {
		println("The answer is 42")
	} else {
		println("The answer is greater than 50")
	}

	if anotherAnswer := -41; anotherAnswer < 0 {
		println("Another answer is less than 0")
	} else if anotherAnswer == 0 {
		println("Another answer is 0")
	} else {
		println("Another answer is greater than 0")
	}
}
