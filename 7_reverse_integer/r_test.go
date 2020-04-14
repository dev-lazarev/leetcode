package __reverse_integer

import "testing"

type testCase struct {
	Input  int
	Output int
}

type testCases struct {
	Tc []testCase
}

func (t *testCases) Add(tc testCase) {
	t.Tc = append(t.Tc, tc)
}

func TestReverse(t *testing.T) {
	tt := testCases{}

	tt.Add(testCase{
		Input:  123,
		Output: 321,
	})

	tt.Add(testCase{
		Input:  -123,
		Output: -321,
	})

	tt.Add(testCase{
		Input:  120,
		Output: 21,
	})

	for _, c := range tt.Tc {
		if result := reverse(c.Input); result != c.Output {
			t.Fatalf("reverse(%v) != %v, expectrd = %v", c.Input, result, c.Output)
		}
	}
}
