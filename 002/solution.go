package s002

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	remainder := 0
	origin := &ListNode{l1.Val + l2.Val, nil}
	l1, l2 = l1.Next, l2.Next
	remainder = origin.Val / 10
	origin.Val %= 10

	var pv *ListNode
	pv = origin
	for l1 != nil || l2 != nil || remainder != 0 {
		nx := &ListNode{remainder, nil}
		if l1 != nil {
			nx.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			nx.Val += l2.Val
			l2 = l2.Next
		}
		remainder = nx.Val / 10
		nx.Val %= 10
		pv.Next = nx
		pv = nx
	}
	return origin
}
