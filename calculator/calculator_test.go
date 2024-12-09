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
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 6, b: 2, want: 4},
		{a: 1, b: 1, want: 0},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Subtract(%f): want %f, got %f", tc.a, tc.b, tc.want)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 1},
		{a: 5, b: 0, want: 0},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Multiply(%f): want %f, got %f", tc.a, tc.b, tc.want)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: 9, b: 3, want: 3},
		{a: 10, b: 2, want: 5},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("Divide(%f, %f): unexpected error: %v", tc.a, tc.b, err)
		}
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Divide(%f): want %f, got %f", tc.a, tc.b, tc.want)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Errorf("You're got an error '%v'", err)
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a    float64
		want float64
	}
	testCases := []testCase{
		{a: 50, want: 7},
		{a: 1, want: 1},
		{a: 25, want: 5},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		if err != nil {
			t.Errorf("Sqrt(%f): unexpected error: %v", tc.a, err)
		}
		if !closeEnough(tc.want, got, 0.1) {
			t.Errorf("Sqrt(%f): want %f, got %f", tc.a, tc.want, got)
		}
	}
}
