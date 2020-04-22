package _2_integer_to_roman

import "testing"

type testCase struct {
	Input  int
	Output string
}

type testCases struct {
	Tc []testCase
}

func (t *testCases) Add(tc testCase) {
	t.Tc = append(t.Tc, tc)
}

func TestIntToRoman(t *testing.T) {

	tt := testCases{}

	tt.Add(testCase{
		Input:  3,
		Output: "III",
	})

	tt.Add(testCase{
		Input:  4,
		Output: "IV",
	})

	tt.Add(testCase{
		Input:  5,
		Output: "V",
	})

	tt.Add(testCase{
		Input:  58,
		Output: "LVIII",
	})

	tt.Add(testCase{
		Input:  1994,
		Output: "MCMXCIV",
	})

	tt.Add(testCase{
		Input:  3999,
		Output: "MMMCMXCIX",
	})

	tt.Add(testCase{
		Input:  99,
		Output: "XCIX",
	})

	for _, c := range tt.Tc {
		if result := intToRoman(c.Input); result != c.Output {
			t.Fatalf("intToRoman(%v) != %v, expectrd = %v", c.Input, result, c.Output)
		}
	}
}
