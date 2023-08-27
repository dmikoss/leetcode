/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	len1 := countLen(l1)
	len2 := countLen(l2)

	if len2 > len1 {
		l1, l2 = l2, l1
	}
	return sumLists(l1, l2)
}

func sumLists(l1 *ListNode, l2 *ListNode) *ListNode {
	iter1 := l1
	iter2 := l2
	overflow := 0
	for iter1 != nil {
		if iter2 != nil {
			fmt.Printf("%v + %v \n", iter1.Val, iter2.Val)
			if iter1.Val+iter2.Val+overflow > 9 {
				iter1.Val = (iter1.Val + iter2.Val + overflow) % 10
				overflow = 1
			} else {
				iter1.Val = iter1.Val + iter2.Val + overflow
				overflow = 0
			}
			iter2 = iter2.Next
		} else {
			if iter1.Val+overflow > 9 {
				iter1.Val = (iter1.Val + overflow) % 10
				overflow = 1
			} else {
				iter1.Val = iter1.Val + overflow
				overflow = 0
			}
		}
		if iter1.Next == nil && overflow > 0 {
			iter1.Next = &ListNode{Val: 1, Next: nil}
			break
		}
		iter1 = iter1.Next
	}
	return l1
}

func countLen(list *ListNode) int {
	count := 1
	for list.Next != nil {
		count++
		list = list.Next
	}
	return count
}
