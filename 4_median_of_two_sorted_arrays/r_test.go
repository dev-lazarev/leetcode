package __median_of_two_sorted_arrays

import "testing"

type testCase struct {
	Input1 []int
	Input2 []int
	Output float64
}

type testCases struct {
	Tc []testCase
}

func (t *testCases) Add(tc testCase) {
	t.Tc = append(t.Tc, tc)
}

func TestFindMedianSortedArrays(t *testing.T) {
	tt := testCases{}
	tt.Add(testCase{
		Input1: []int{1, 3},
		Input2: []int{2},
		Output: 2.0,
	})
	tt.Add(testCase{
		Input1: []int{1, 2},
		Input2: []int{3, 4},
		Output: 2.5,
	})

	for _, c := range tt.Tc {
		if result := findMedianSortedArrays(c.Input1, c.Input2); result != c.Output {
			t.Fatalf("findMedianSortedArrays(%v, %v) != %f, expectrd = %f", c.Input1, c.Input2, result, c.Output)
		}
	}
}
