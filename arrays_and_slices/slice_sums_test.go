package list

import (
	"reflect"
	"testing"
)

func TestSliceSums(t *testing.T) {
	got := SliceSums([]int{3, 4, 5}, []int{6, 7, 8})
	want := []int{12, 21}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got: %v, Want: %v", got, want)
	}
}
