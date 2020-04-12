package __median_of_two_sorted_arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result float64
	merged := merge(nums1, nums2)

	l := len(merged)

	if l%2 == 1 {
		index := (l - 1) / 2
		result = float64(merged[index])
	} else {
		index := l / 2
		l1 := merged[index-1]
		l2 := merged[index]
		result = float64(l1) + float64(l2)
		result = result / 2
	}
	return result
}

func merge(nums1 []int, nums2 []int) []int {
	lenA := len(nums1)
	lenB := len(nums2)

	i := lenA - 1
	j := lenB - 1

	result := make([]int, lenA+lenB, lenA+lenB)
	copy(result, nums1)

	end := lenA + lenB - 1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			result[end] = nums1[i]
			i--
		} else {
			result[end] = nums2[j]
			j--
		}
		end--
	}
	return result
}
