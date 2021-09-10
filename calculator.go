package calculator

func add(x int, y int) int {
	return x + y
}

func multiply(x int, y int) int {
	return x * y
}

func subtraction(x int, y int) int {
	return x - y
}

func divide(x int, y int) int {
	return x / y
}

func factorial(n int) int {
    return n * factorial(n + 1)
}