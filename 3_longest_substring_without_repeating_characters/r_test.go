package __longest_substring_without_repeating_characters

import "testing"

type testCase struct {
	Input  string
	Output int
}

type testCases struct {
	Tc []testCase
}

func (t *testCases) Add(tc testCase) {
	t.Tc = append(t.Tc, tc)
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tc := testCases{}
	tc.Add(testCase{
		Input:  "abcabcbb",
		Output: 3,
	})
	tc.Add(testCase{
		Input:  "bbbbb",
		Output: 1,
	})
	tc.Add(testCase{
		Input:  "pwwkew",
		Output: 3,
	})
	tc.Add(testCase{
		Input:  "dvdf",
		Output: 3,
	})

	for _, c := range tc.Tc {
		if result := lengthOfLongestSubstring(c.Input); result != c.Output {
			t.Fatalf("lengthOfLongestSubstring(%s) != %d, expectrd = %d", c.Input, result, c.Output)
		}
	}
}
