package calculator

import "testing"

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


	got := multiply(3, 3)
	if got != 9 {
		t.Error("Calculator.Multiply(3, 3) =", got, "; should", 3*3)
	}
}

func TestSubstraction(t *testing.T) {
	got := subtraction(3, 3)
	if got != 0 {
		t.Error("Calculator.Substraction(3, 3) =", got, "; should", 3-3)
	}
}

func TestDivide(t *testing.T) {
	got := divide(3, 3)
	if got != 1 {
		t.Error("Calculator.Divide(3, 3) =", got, "; should", 3/3)
	}
}
