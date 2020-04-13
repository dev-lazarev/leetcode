package __zigzag_conversion

import "testing"

type testCase struct {
	Input1 string
	Input2 int
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
		Input1: "PAYPALISHIRING",
		Input2: 3,
		Output: "PAHNAPLSIIGYIR",
	})

	tt.Add(testCase{
		Input1: "PAYPALISHIRING",
		Input2: 2,
		Output: "PYAIHRNAPLSIIG",
	})

	tt.Add(testCase{
		Input1: "PAYPALISHIRING",
		Input2: 20,
		Output: "PAYPALISHIRING",
	})

	tt.Add(testCase{
		Input1: "PAYPALISHIRING",
		Input2: 4,
		Output: "PINALSIGYAHRPI",
	})

	tt.Add(testCase{
		Input1: "PAYPALISHIRING",
		Input2: 5,
		Output: "PHASIYIRPLIGAN",
	})

	for _, c := range tt.Tc {
		if result := convert(c.Input1, c.Input2); result != c.Output {
			t.Fatalf("convert(%v, %v) != %v, expectrd = %v", c.Input1, c.Input2, result, c.Output)
		}
	}
}
