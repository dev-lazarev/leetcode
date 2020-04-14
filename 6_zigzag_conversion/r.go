package __zigzag_conversion

func convert(s string, numRows int) string {
	result := ""
	length := len(s)
	if length <= numRows {
		return s
	}

	array := make([]string, numRows)

	for index, row := 0, 0; index < length; {
		if row < numRows {
			array[row] += string(s[index])
			row++
			index++
			continue
		}

		for subRow := numRows - 1; subRow > 0 && index < length; subRow-- {
			if subRow == numRows-1 {
				continue
			}
			array[subRow] += string(s[index])
			index++
		}
		row = 0
	}

	for _, substring := range array {
		result += substring
	}
	return result
}
