package __two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	input1 := []int{2, 7, 11, 15}
	input2 := 9
	expected := []int{0, 1}
	if result := twoSum(input1, input2); !reflect.DeepEqual(twoSum(input1, input2), expected) {
		t.Fatalf("twoSum(%v, %v) != %v, expectrd = %v", input1, input2, result, expected)
	}
}
