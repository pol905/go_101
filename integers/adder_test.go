package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4

	if got != want {
		t.Errorf("Got: %d, Want: %d", got, want)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
