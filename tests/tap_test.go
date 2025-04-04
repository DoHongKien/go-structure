package tests

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func Add(i1, i2 int) int {
	return i1 + i2
}
