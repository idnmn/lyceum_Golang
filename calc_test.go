package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		expression string
		want       float64
	}{
		{"(1+1)*2", 4},
		{"2*(1+1)", 4},
		{"3*(2+2*3)", 24},
		{"(2*(2+3))/5", 2},
		{"(2*(2+3) + 5)/5", 3},
		{"10.5-20", -9.5},
		{"2+3*2", 8},
		{"sdafasdf", 0},
		{"2+asd", 0},
		{"99--1", 0},
	}

	for _, test := range tests {
		got, err := Calc(test.expression)
		if err != nil {
			t.Errorf("Calc(%q) returned error: %v", test.expression, err)
		} else if got != test.want {
			t.Errorf("Calc(%q) = %f, want %f", test.expression, got, test.want)
		}
	}
}
