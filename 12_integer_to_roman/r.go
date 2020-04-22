package _2_integer_to_roman

var table = map[int]map[int]string{
	0: {1: "I", 5: "V"},
	1: {1: "X", 5: "L"},
	2: {1: "C", 5: "D"},
	3: {1: "M", 5: ""},
}

func intToRoman(num int) string {
	if num < 1 || num > 3999 {
		return ""
	}

	result := ""
	for i := 0; num > 0; i++ {
		r := ""
		x := num % 10
		num = (num - x) / 10
		if x == 9 {
			r = table[i][1] + table[i+1][1]
		} else if x < 9 && x >= 5 {
			r = table[i][5]
			for x = x - 5; x > 0; x-- {
				r = r + table[i][1]
			}
		} else if x == 4 {
			r = table[i][1] + table[i][5]
		} else {
			for ; x > 0; x-- {
				r = r + table[i][1]
			}
		}
		result = r + result
	}
	return result
}
