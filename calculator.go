package calculator

func add(x int, y int) int {
	return x + y
}

func multiply(x int, y int) int {
	return x * y
}

func subtraction(x int, y int) int {
	//self-assigned variable
	x = x

	//duplicated string literals
	var a string = "Hello"
	var b string = "Hello"
	var c string = "Hello"
	if a == "Hello" && b == "Hello" && c == "Hello" {
	return x - y
	} else if a == "Hello" && b == "Hello" && c == "Hello" { //same condition
	return x - y
	} else {
	return 1
	}
}

//empty function
func doNothing() {
}

func divide(x int, y int) int {
	return x / y
}
