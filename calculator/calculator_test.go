package calculator_test

import (
	"testing"

	"github.com/athunlal/calculator"
)

var TestCase = []struct {
	name     string
	divident int
	diviser  int
	expected int
}{
	{"division", 10, 5, 2},
	{"division_negative", -50, 10, -5},
}

func TestDive(t *testing.T) {

	for _, tc := range TestCase {

		t.Run(tc.name, func(t *testing.T) {

			if got := calculator.Divide(tc.divident, tc.diviser); got != tc.expected {
				t.Errorf("expected %d got %d ", tc.expected, got)
			}

		})

	}
}

var TestCase2 = []struct {
	name        string
	firstValue  int
	secondValue int
	expected    int
}{
	{"addition", 10, 5, 15},
	{"addition_negative", -50, 10, -40},
}

func TestAdd(t *testing.T) {
	for _, tc := range TestCase2 {
		t.Run(tc.name, func(t *testing.T) {
			if got := calculator.Add(tc.firstValue, tc.secondValue); got != tc.expected {
				t.Errorf("expected %d got %d", tc.expected, got)
			}
		})
	}
}

var TestCase3 = []struct {
	name        string
	firstValue  int
	secondValue int
	expected    int
}{
	{"subtract", 10, 5, 5},
	{"subtract_negative", 5, 10, -5},
}

func TestSubstract(t *testing.T) {
	for _, tc := range TestCase3 {
		t.Run(tc.name, func(t *testing.T) {
			if got := calculator.Substract(tc.firstValue, tc.secondValue); got != tc.expected {
				t.Errorf("expected %d got %d", tc.expected, got)
			}
		})
	}
}

var TestCase4 = []struct{
	name string
	firstValue int
	secondValue int
	expected int
}{
	{"multiplication", 10, 5, 50},
	{"multiplication_nagative", 10, -5, -50},
}

func TestMultiplication(t *testing.T){
	for _,tc := range TestCase4{
		t.Run(tc.name, func(t *testing.T) {
			if got := calculator.Multiplication(tc.firstValue, tc.secondValue); got != tc.expected{
				t.Errorf("expected %d got %d", tc.expected, got)
			}
		})
	}
}