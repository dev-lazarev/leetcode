package __add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := ListNodeToInt(l1)
	n2 := ListNodeToInt(l2)
	max, min := n2, n1
	if len(n1) > len(n2) {
		max, min = n1, n2
	}

	result := make([]int, 0)
	maxLen := len(max)
	minLen := len(min)
	needAppend := false
	for i := 0; i <= maxLen; i++ {
		if i == maxLen && needAppend == false {
			break
		}
		sum := 0
		if i < maxLen {
			sum += max[i]
		}
		if i < minLen {
			sum += min[i]
		}
		if needAppend == true {
			sum++
		}
		if sum > 9 {
			sum = sum - 10
			needAppend = true
		} else {
			needAppend = false
		}
		result = append(result, sum)
	}
	return createNode(result)

}

func ListNodeToInt(l *ListNode) []int {
	result := make([]int, 0)
	if l == nil {
		return result
	}
	result = append(result, l.Val)
	result = append(result, ListNodeToInt(l.Next)...)
	return result
}

func createNode(n []int) *ListNode {
	node := ListNode{}
	if len(n) > 0 {
		node.Val = n[0]
		node.Next = createNode(n[1:])
		return &node
	} else {
		return nil
	}
}
