package __longest_palindromic_substring

func longestPalindrome(s string) string {
	length := len(s)
	if length == 0 {
		return ""
	}

	result := s[:1]
	lastLength := 1

	for indexMiddle := 0; indexMiddle < length; indexMiddle++ {
		leftBorder, rightBorder := indexMiddle, indexMiddle+1
		for leftBorder >= 0 && rightBorder < length && s[leftBorder] == s[rightBorder] {
			currentLength := rightBorder - leftBorder + 1
			if currentLength > lastLength {
				lastLength = currentLength
				result = s[leftBorder : rightBorder+1]
			}
			leftBorder--
			rightBorder++
		}
	}

	for indexMiddle := 0; indexMiddle < length; indexMiddle++ {
		leftBorder, rightBorder := indexMiddle-1, indexMiddle+1
		for leftBorder >= 0 && rightBorder < length && s[leftBorder] == s[rightBorder] {
			currentLength := rightBorder - leftBorder + 1
			if currentLength > lastLength {
				lastLength = currentLength
				result = s[leftBorder : rightBorder+1]
			}
			leftBorder--
			rightBorder++
		}
	}

	return result
}
