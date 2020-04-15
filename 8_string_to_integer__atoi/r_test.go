package __string_to_integer__atoi

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

func TestMyAtoi(t *testing.T) {
	tt := testCases{}

	tt.Add(testCase{
		Input:  " -01234567890",
		Output: -1234567890,
	})

	tt.Add(testCase{
		Input:  "0",
		Output: 0,
	})

	tt.Add(testCase{
		Input:  "-123",
		Output: -123,
	})

	tt.Add(testCase{
		Input:  "asd123",
		Output: 0,
	})

	tt.Add(testCase{
		Input:  "   300005",
		Output: 300005,
	})

	tt.Add(testCase{
		Input:  "   300005  333",
		Output: 300005,
	})

	tt.Add(testCase{
		Input:  "4193 with words",
		Output: 4193,
	})

	tt.Add(testCase{
		Input:  "words and 987",
		Output: 0,
	})

	tt.Add(testCase{
		Input:  "-91283472332",
		Output: -2147483648,
	})

	tt.Add(testCase{
		Input:  "+1",
		Output: 1,
	})

	for _, c := range tt.Tc {
		if result := myAtoi(c.Input); result != c.Output {
			t.Fatalf("myAtoi(%v) != %v, expectrd = %v", c.Input, result, c.Output)
		}
	}
}
