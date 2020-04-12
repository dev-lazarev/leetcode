package __longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	su := make(map[uint8]int)
	max := 0

	for i := 0; i < len(s); i++ {
		w := s[i]
		if ind, ok := su[w]; ok {
			max = max2(su, max)
			su = make(map[uint8]int)
			i = ind
		} else {
			su[w] = i
		}
	}

	return max2(su, max)
}

func max2(su map[uint8]int, max int) int {
	l := len(su)
	if l > max {
		max = l
	}
	return max
}
