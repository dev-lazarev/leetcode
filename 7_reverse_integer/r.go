package __reverse_integer

const MaxInt32 = 1<<31 - 1
const MinInt32 = -1 << 31

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	result := 0
	factor := 1
	if x < 0 {
		factor = -1
		x *= factor
	}

	ns := numSize(x)

	for i := 1; i <= ns; i++ {
		n := x % 10
		x = (x - n) / 10
		if result == 0 && n == 0 {
			continue
		}
		result += n * pow10(ns-i)
	}

	result *= factor

	if result > MaxInt32 || result < MinInt32 {
		return 0
	}
	return result
}

func pow10(x int) int {
	if x == 0 {
		return 1
	}
	p := 10
	for i := 1; i < x; i++ {
		p = 10 * p
	}
	return p
}

func numSize(x int) int {
	p := 10
	if x < 0 {
		x = x * -1
	}
	for i := 1; i < 10; i++ {
		if x < p {
			return i
		}
		p = 10 * p
	}
	return 10
}
