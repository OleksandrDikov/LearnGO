package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 1, b: 2, want: 3},
		{a: 3, b: 4, want: 7},
		{a: 4, b: 5, want: 9},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 1, b: 2, want: -1},
		{a: 7, b: 4, want: 3},
		{a: 12, b: 5, want: 7},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 1, b: 2, want: 2},
		{a: 3, b: 4, want: 12},
		{a: 4, b: 5, want: 20},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 4, b: 2, want: 2},
		{a: 16, b: 4, want: 4},
		{a: 25, b: 5, want: 5},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Error("want error for invalid input, got nil")
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 25, want: 5},
		{a: 9, want: 3},
		{a: 64, want: 8},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if !closeEnough(tc.want, got, 0.1) {
			t.Errorf("Sqrt(%f): want %f, got %f", tc.a, tc.want, got)
		}
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
