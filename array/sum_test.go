package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	actual := Sum(numbers)
	expected := 15

	if actual != expected {
		t.Errorf("not equal actual %q expected %q", actual, expected)
	}
}

func TestSumAll(t *testing.T) {
	actual := SumAll([]int{1, 2, 3}, []int{4, 5})
	expected := []int{6, 9}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("not equal actual %q expected %q", actual, expected)
	}
}
