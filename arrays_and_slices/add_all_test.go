package list

import "testing"

func TestAddAll(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := AddAll(nums)
	expected := 55

	if got != expected {
		t.Errorf("Got: %d, Want: %d", got, expected)
	}
}
