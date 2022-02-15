package calculator

import "testing"

func TestAdd(t *testing.T) {
	//Student test repository
	got := add(3, 3)
	if got != 6 {
		t.Error("Calculator.Add(3, 3) =", got, "; should", 3+3)
	}
}
