package __palindrome_number

import "testing"

type testCase struct {
	Input  int
	Output bool
}

type testCases struct {
	Tc []testCase
}

func (t *testCases) Add(tc testCase) {
	t.Tc = append(t.Tc, tc)
}

func TestIsPalindrome(t *testing.T) {
	tt := testCases{}

	tt.Add(testCase{
		Input:  121,
		Output: true,
	})

	tt.Add(testCase{
		Input:  -121,
		Output: false,
	})

	tt.Add(testCase{
		Input:  1331,
		Output: true,
	})

	tt.Add(testCase{
		Input:  10,
		Output: false,
	})

	tt.Add(testCase{
		Input:  0,
		Output: true,
	})

	tt.Add(testCase{
		Input:  9,
		Output: true,
	})

	tt.Add(testCase{
		Input:  23,
		Output: false,
	})

	for _, c := range tt.Tc {
		if result := isPalindrome(c.Input); result != c.Output {
			t.Fatalf("isPalindrome(%v) != %v, expectrd = %v", c.Input, result, c.Output)
		}
	}
}
