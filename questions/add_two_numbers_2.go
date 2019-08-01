package questions

import (
	"fmt"
	"reflect"

	track "github.com/OscarZhou/gotrack"
)

type AddTwoNumbers2 struct {
	Question
	Track *track.Track
}

func (e *AddTwoNumbers2) Init() {
	e.No = 445
	e.Title = "Add Two Numbers II"
	e.FullTitle = "Add Two Numbers II"
	e.URL = "https://leetcode.com/problems/add-two-numbers-ii/"
	e.Level = LevelMedium
	e.FuncName = "addTwoNumbers2"
	NumberMap[e.No] = e
	TitleMap[e.Title] = e
}

func (e AddTwoNumbers2) PrintTitle() {
	fmt.Printf("%d, %s\n", e.No, e.Title)
}

func (e AddTwoNumbers2) PrintDetail() {
	fmt.Printf("%d, %s\n", e.No, e.Title)
}

func (e AddTwoNumbers2) Run() error {
	e.Track.Start()
	defer e.Track.End()

	l1 := &ListNode{Val: 7, Next: nil}
	header1 := l1
	l1.Next = &ListNode{Val: 2, Next: nil}
	l1 = l1.Next
	l1.Next = &ListNode{Val: 4, Next: nil}
	l1 = l1.Next
	l1.Next = &ListNode{Val: 3, Next: nil}
	l1 = l1.Next
	l1.Next = nil

	l2 := &ListNode{Val: 5, Next: nil}
	header2 := l2
	l2.Next = &ListNode{Val: 6, Next: nil}
	l2 = l2.Next
	l2.Next = &ListNode{Val: 4, Next: nil}
	l2 = l2.Next
	l2.Next = nil

	expect := &ListNode{Val: 7, Next: nil}
	headerExpect := expect
	expect.Next = &ListNode{Val: 8, Next: nil}
	expect = expect.Next
	expect.Next = &ListNode{Val: 0, Next: nil}
	expect = expect.Next
	expect.Next = &ListNode{Val: 7, Next: nil}
	expect = expect.Next
	expect.Next = nil

	actual := addTwoNumbers2(header1, header2)
	if !reflect.DeepEqual(*headerExpect, *actual) {
		fmt.Printf("expected=%v, actual=%v\n", headerExpect, actual)
	}

	// [5],[5]
	// [1,9], [9]
	return nil
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var headNode, cur *ListNode
	len1, len2 := 0, 0
	//get length
	for cur1 := l1; cur1 != nil; cur1 = cur1.Next {
		len1++
	}
	for cur2 := l2; cur2 != nil; cur2 = cur2.Next {
		len2++
	}

	if len1 == len2 && len2 == 0 {
		return nil
	}

	//align with l1,l2
	if len1 < len2 {
		len1, len2 = len2, len1
		l1, l2 = l2, l1
	}

	if len1 != len2 {
		for i := 0; i < len1-len2; i++ {
			if i == 0 {
				headNode = new(ListNode)
				cur = headNode
			} else {
				cur.Next = new(ListNode)
				cur = cur.Next
			}
		}
		cur.Next = l2
		l2 = headNode
	}

	// recursion
	carry, headNode := addon(l1, l2, 0)
	if carry > 0 {
		temp := &ListNode{
			Val:  carry,
			Next: headNode,
		}
		headNode = temp
	}

	return headNode
}

func addon(cur1 *ListNode, cur2 *ListNode, carry int) (int, *ListNode) {
	var node *ListNode
	if cur1.Next != nil {
		carry, node = addon(cur1.Next, cur2.Next, carry)
	}
	newNode := &ListNode{
		Val: (cur1.Val + cur2.Val + carry) % 10,
	}
	newNode.Next = node
	return (cur1.Val + cur2.Val + carry) / 10, newNode
}

func (e AddTwoNumbers2) Print() {
	fmt.Println(`
	func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
		var headNode, cur *ListNode
		len1, len2 := 0, 0
		//get length
		for cur1 := l1; cur1 != nil; cur1 = cur1.Next {
			len1++
		}
		for cur2 := l2; cur2 != nil; cur2 = cur2.Next {
			len2++
		}
	
		if len1 == len2 && len2 == 0 {
			return nil
		}
	
		//align with l1,l2
		if len1 < len2 {
			len1, len2 = len2, len1
			l1, l2 = l2, l1
		}
	
		if len1 != len2 {
			for i := 0; i < len1-len2; i++ {
				if i == 0 {
					headNode = new(ListNode)
					cur = headNode
				} else {
					cur.Next = new(ListNode)
					cur = cur.Next
				}
			}
			cur.Next = l2
			l2 = headNode
		}
	
		// recursion
		carry, headNode := addon(l1, l2, 0)
		if carry > 0 {
			temp := &ListNode{
				Val:  carry,
				Next: headNode,
			}
			headNode = temp
		}
	
		return headNode
	}
	
	func addon(cur1 *ListNode, cur2 *ListNode, carry int) (int, *ListNode) {
		var node *ListNode
		if cur1.Next != nil {
			carry, node = addon(cur1.Next, cur2.Next, carry)
		}
		newNode := &ListNode{
			Val: (cur1.Val + cur2.Val + carry) % 10,
		}
		newNode.Next = node
		return (cur1.Val + cur2.Val + carry) / 10, newNode
	}
	
	`)

}
