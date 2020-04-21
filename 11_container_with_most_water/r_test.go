package _1_container_with_most_water

import "testing"

type testCase struct {
	Input  []int
	Output int
}

type testCases struct {
	Tc []testCase
}

func (t *testCases) Add(tc testCase) {
	t.Tc = append(t.Tc, tc)
}

func TestMaxArea(t *testing.T) {
	tt := testCases{}

	tt.Add(testCase{
		Input:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		Output: 49,
	})

	tt.Add(testCase{
		Input:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		Output: 0,
	})

	tt.Add(testCase{
		Input:  []int{0, 1, 2, 3},
		Output: 2,
	})

	tt.Add(testCase{
		Input:  []int{3, 3, 3, 3},
		Output: 9,
	})

	tt.Add(testCase{
		Input:  []int{3, 3, 3, 3, 3},
		Output: 12,
	})

	tt.Add(testCase{
		Input:  []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		Output: 12,
	})

	tt.Add(testCase{
		Input:  []int{2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		Output: 12,
	})

	tt.Add(testCase{
		Input:  []int{1},
		Output: 0,
	})

	tt.Add(testCase{
		Input:  []int{},
		Output: 0,
	})

	tt.Add(testCase{
		Input:  []int{1, 1},
		Output: 1,
	})

	for _, c := range tt.Tc {
		if result := maxArea(c.Input); result != c.Output {
			t.Fatalf("maxArea(%v) != %v, expectrd = %v", c.Input, result, c.Output)
		}
	}
}
