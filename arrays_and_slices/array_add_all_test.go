package list

import "testing"

func TestArrayAddAll(t *testing.T) {
	got := ArrayAddAll([5]int{1, 2, 3, 4, 5})
	want := 15

	if got != want {
		t.Errorf("Got: %q, want: %q", got, want)
	}

}
