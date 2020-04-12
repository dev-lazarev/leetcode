package __add_two_numbers

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	n1 := []int{9}
	n2 := []int{9}

	l1 := createNode(n1)
	l2 := createNode(n2)

	expected := []int{8, 1}

	if result := ListNodeToInt(addTwoNumbers(l1, l2)); !reflect.DeepEqual(result, expected) {
		t.Fatalf("addTwoNumbers(%v, %v) != %v, expected = %v", n1, n2, result, expected)
	}
}
