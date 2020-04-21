package _1_container_with_most_water

func maxArea(height []int) int {
	h := len(height)
	if h < 2 {
		return 0
	}

	result := 0
	for i := 0; i < h-1; i++ {
		for j := h - 1; j > i && result <= height[i]*(j-1); j-- {
			l := (j - i) * min(height[i], height[j])
			result = max(result, l)
		}
	}

	return result
}

func min(n1 int, n2 int) int {
	if n1 > n2 {
		return n2
	} else {
		return n1
	}
}

func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}
