package test

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var pRetListNode *ListNode = new(ListNode)
	var pRoot *ListNode = pRetListNode
	var carry int = 0

	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		pRetListNode.Val = sum % 10
		if l1 != nil || l2 != nil {
			pRetListNode.Next = new(ListNode)
			pRetListNode = pRetListNode.Next
		}
	}
	if 0 < carry {
		pRetListNode.Next = new(ListNode)
		pRetListNode = pRetListNode.Next
		pRetListNode.Val = 1
	}
	return pRoot
}

func MakeList(arr []int) *ListNode {
	var pRetListNode *ListNode = new(ListNode)
	var pRoot *ListNode = pRetListNode

	for i, v := range arr {
		pRetListNode.Val = v
		if i < len(arr)-1 {
			pRetListNode.Next = new(ListNode)
			pRetListNode = pRetListNode.Next
		}
	}

	return pRoot
}

func TestFunction(a *int) {
	*a = 2
}
