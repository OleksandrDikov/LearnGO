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
