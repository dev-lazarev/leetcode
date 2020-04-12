package __two_sum

func twoSum(nums []int, target int) []int {
	var a []int
	for index, num := range nums {
		for index2, num2 := range nums {
			if index == index2 {
				continue
			}

			if num+num2 == target {
				a = append(a, index, index2)
				return a
			}
		}
	}
	return a
}
