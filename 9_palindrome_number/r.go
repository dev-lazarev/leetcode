package __palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	if x%10 == 0 {
		return false
	}
	a := make([]int, 0)

	i := 0
	for ; x != 0; i++ {
		sx := x % 10
		x = x - sx
		a = append(a, sx)
		x = x / 10
	}

	for lb, rb := 0, i-1; lb < rb; lb, rb = lb+1, rb-1 {
		if a[lb] != a[rb] {
			return false
		}
	}

	return true
}
