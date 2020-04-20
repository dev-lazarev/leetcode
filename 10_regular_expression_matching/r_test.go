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

type testCase2 struct {
	Input1 string
	Input2 uint
	Output string
}

type testCases2 struct {
	Tc []testCase2
}

func (t *testCases2) Add(tc testCase2) {
	t.Tc = append(t.Tc, tc)
}

func TestSub(t *testing.T) {
	tc := testCases2{}
	tc.Add(testCase2{
		Input1: "aa",
		Input2: 1,
		Output: "a",
	})
	tc.Add(testCase2{
		Input1: "aa",
		Input2: 0,
		Output: "aa",
	})
	tc.Add(testCase2{
		Input1: "aa",
		Input2: 2,
		Output: "",
	})
	tc.Add(testCase2{
		Input1: "aa",
		Input2: 3,
		Output: "",
	})

	tc.Add(testCase2{
		Input1: "ab",
		Input2: 1,
		Output: "b",
	})
	tc.Add(testCase2{
		Input1: "abc",
		Input2: 1,
		Output: "bc",
	})

	for _, c := range tc.Tc {
		if result := sub(c.Input1, c.Input2); result != c.Output {
			t.Fatalf("sub(%v, %v) != %v, expectrd = %v", c.Input1, c.Input2, result, c.Output)
		}
	}
}

func TestIsMatch(t *testing.T) {
	tc := testCases{}
	tc.Add(testCase{
		Input1: "a",
		Input2: ".*..a*",
		Output: false,
	})
	tc.Add(testCase{
		Input1: "bbbba",
		Input2: ".*a*a",
		Output: true,
	})
	tc.Add(testCase{
		Input1: "a",
		Input2: "ab*",
		Output: true,
	})
	tc.Add(testCase{
		Input1: "i",
		Input2: ".",
		Output: true,
	})
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
	tc.Add(testCase{
		Input1: "mississippi",
		Input2: "mis*is*ip*.",
		Output: true,
	})

	tc.Add(testCase{
		Input1: "aaa",
		Input2: "a*a",
		Output: true,
	})
	tc.Add(testCase{
		Input1: "aaa",
		Input2: "ab*a*c*a",
		Output: true,
	})

	for _, c := range tc.Tc {
		if result := isMatch(c.Input1, c.Input2); result != c.Output {
			t.Fatalf("isMatch(%v, %v) != %v, expectrd = %v", c.Input1, c.Input2, result, c.Output)
		}
	}
}
