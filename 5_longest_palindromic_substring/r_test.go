package __longest_palindromic_substring

import "testing"

type testCase struct {
	Input  string
	Output string
}

type testCases struct {
	Tc []testCase
}

func (t *testCases) Add(tc testCase) {
	t.Tc = append(t.Tc, tc)
}

func TestLongestPalindrome(t *testing.T) {

	tt := testCases{}

	tt.Add(testCase{
		Input:  "babad",
		Output: "bab",
	})

	tt.Add(testCase{
		Input:  "cbbd",
		Output: "bb",
	})

	tt.Add(testCase{
		Input:  "a",
		Output: "a",
	})

	tt.Add(testCase{
		Input:  "ac",
		Output: "a",
	})

	tt.Add(testCase{
		Input:  "abb",
		Output: "bb",
	})

	for _, c := range tt.Tc {
		if result := longestPalindrome(c.Input); result != c.Output {
			t.Fatalf("longestPalindrome(%v) != %v, expectrd = %v", c.Input, result, c.Output)
		}
	}
}
