package _0_regular_expression_matching

func isMatch(s string, p string) bool {
	lengthP := len(p)
	lengthS := len(s)
	if lengthP == 0 {
		return lengthS == 0
	}
	s0 := uint8(0)
	if lengthS > 0 {
		s0 = s[0]
	}

	//*
	if p[0] == 42 {
		return false
	}
	result := false

	//a-z || .
	if p[0] >= 97 && p[0] <= 122 || p[0] == 46 {
		if s0 == p[0] || p[0] == 46 {
			if lengthP > 1 && p[1] == 42 {
				if lengthS == 0 && lengthP > 1 {
					result = isMatch(sub(s, 1), sub(p, 2))
				} else {
					result = isMatch(sub(s, 1), sub(p, 0)) ||
						isMatch(sub(s, 0), sub(p, 2))
				}
			} else {
				result = s0 != 0 && isMatch(sub(s, 1), sub(p, 1))
			}
		} else {
			if lengthP > 1 && p[1] == 42 {
				result = isMatch(sub(s, 0), sub(p, 2))
			}
		}
	}

	return result
}

func sub(s string, leftIndex uint) string {
	if len(s) < int(leftIndex)+1 {
		return ""
	} else {
		return s[leftIndex:]
	}
}
