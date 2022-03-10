package calculator

import "testing"
import "fmt"

func TestAdd(t *testing.T) {
	got := add(3, 3)
	if got != 6 {
		t.Error("Calculator.Add(3, 3) =", got, "; should", 3+3)
	}
}

func TestMultiply(t *testing.T) {
	//duplicated string literals
	var a string = "Hello"
	var b string = "Hello"
	var c string = "Hello"
	if a == "Hello" && b == "Hello" && c == "Hello" {
	fmt.Println("a is: ", a)
		} else {
	fmt.Println("b is: ", b)
	}



	got := multiply(3, 3)
	if got != 9 {
		t.Error("Calculator.Multiply(3, 3) =", got, "; should", 3*3)
	}
}

func TestDivide(t *testing.T) {
	got := divide(3, 3)
	if got != 1 {
		t.Error("Calculator.Divide(3, 3) =", got, "; should", 3/3)
	}
}
