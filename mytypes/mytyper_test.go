package mytypes_test

import (
	"mytypes"
	"testing"
)

func TestMyInt_Twice(t *testing.T) {
	t.Parallel()
	i := mytypes.MyInt(5)
	want := mytypes.MyInt(10)
	got := i.Twice()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMyStringLen(t *testing.T) {
	t.Parallel()
	s := mytypes.MyString("hello")
	want := 5
	got := s.Len()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
