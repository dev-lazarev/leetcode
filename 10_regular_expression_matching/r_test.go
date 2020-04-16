package _0_regular_expression_matching

import "testing"

type testCase struct {
	Input1 string
	Input2 string
	Output bool
}

type testCases struct {
	Tc []testCase
}

func (t *testCases) Add(tc testCase) {
	t.Tc = append(t.Tc, tc)
}

func TestIsMatch(t *testing.T) {
	tc := testCases{}
	tc.Add(testCase{
		Input1: "aa",
		Input2: "a",
		Output: false,
	})
	tc.Add(testCase{
		Input1: "aa",
		Input2: "a*",
		Output: true,
	})
	tc.Add(testCase{
		Input1: "ab",
		Input2: ".*",
		Output: true,
	})
	tc.Add(testCase{
		Input1: "aab",
		Input2: "c*a*b",
		Output: true,
	})
	tc.Add(testCase{
		Input1: "mississippi",
		Input2: "mis*is*p*.",
		Output: false,
	})

	for _, c := range tc.Tc {
		if result := isMatch(c.Input1, c.Input2); result != c.Output {
			t.Fatalf("isMatch(%v,%v) != %v, expectrd = %v", c.Input1, c.Input2, result, c.Output)
		}
	}
}
