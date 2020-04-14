package __reverse_integer

func reverse(x int) int {
	const MaxInt32 = 1<<31 - 1
	const MinInt32 = -1 << 31
	result := 0
	if result > MaxInt32 || result < MinInt32 {
		return 0
	}
	return result
}
