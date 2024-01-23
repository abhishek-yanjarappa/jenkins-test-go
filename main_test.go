package main
import "testing"

func TestSum(t *testing.T) {
	result := Sum(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Sum(2, 3) expected %d, but got %d", expected, result)
	}
}

func TestSumNegative(t *testing.T) {
	result := Sum(-1, 1)
	expected := 0

	if result != expected {
		t.Errorf("Sum(-1, 1) expected %d, but got %d", expected, result)
	}
}