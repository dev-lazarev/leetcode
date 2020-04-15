package __string_to_integer__atoi

const MaxInt32 = 1<<31 - 1
const MinInt32 = -1 << 31

func myAtoi(str string) int {

	result := 0
	start := false
	factor := 1
	for _, w := range str {
		if w == 32 {
			if !start {
				continue
			} else {
				break
			}
		} else if w == 43 {
			if start {
				break
			} else {
				start = true
			}
		} else if w == 45 {
			if start {
				break
			} else {
				factor = -1
				start = true
			}
		} else if w >= 48 && w <= 57 {
			start = true
			result = result*10 + int(w) - 48
			if result > MaxInt32 {
				break
			}
		} else {
			break
		}

	}
	result *= factor

	if result > MaxInt32 {
		result = MaxInt32
	}
	if result < MinInt32 {
		result = MinInt32
	}
	return result
}
