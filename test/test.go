package test

import (
	"fmt"
	"strings"
	"sync"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func TwoSum(nums []int, target int) []int {
	var prevMap map[int]int = make(map[int]int)
	for i, num := range nums {
		diff := target - num
		value, ok := prevMap[diff]
		if ok {
			return []int{value, i}
		}
		prevMap[num] = i
	}

	return []int{0, 0}
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

func LengthOfLongestSubstring(s string) int {

	var max int
	var t string = ""
	for _, v := range s {
		pos := strings.Index(t, string(v))
		if pos == -1 {
			t += string(v)
			if max < len(t) {
				max = len(t)
			}
		} else {
			if pos < len(t)-1 {
				t = t[pos+1:] + string(v)
			} else {
				t = string(v)
			}
		}
	}

	return max
}

func TestGoRoutine1(ch chan string, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	str := "hi!"
	ch <- str
	fmt.Println("This is goRoutine", id, "!", str, "is sent")
}

func TestGoRoutine2(ch chan string, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	str := <-ch
	fmt.Println("This is goRoutine", id, "!", str, "is received!")
}
