package questions

import (
	"fmt"
	"reflect"

	track "github.com/OscarZhou/gotrack"
)

type AddTwoNumbers struct {
	Question
	Track *track.Track
}

func (e *AddTwoNumbers) Init() {
	e.No = 2
	e.Title = "Add Two Numbers"
	e.FullTitle = "Add Two Numbers"
	e.URL = "https://leetcode.com/problems/add-two-numbers/"
	e.Level = LevelMedium
	e.FuncName = "addTwoNumbers"
	NumberMap[e.No] = e
	TitleMap[e.Title] = e
}

func (e AddTwoNumbers) PrintTitle() {
	fmt.Printf("%d, %s\n", e.No, e.Title)
}

func (e AddTwoNumbers) PrintDetail() {
	fmt.Printf("%d, %s\n", e.No, e.Title)
}

func (e AddTwoNumbers) Run() error {
	e.Track.Start()
	defer e.Track.End()

	l1 := &ListNode{Val: 2, Next: nil}
	header1 := l1
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
	expect.Next = &ListNode{Val: 0, Next: nil}
	expect = expect.Next
	expect.Next = &ListNode{Val: 8, Next: nil}
	expect = expect.Next
	expect.Next = nil

	actual := addTwoNumbers(header1, header2)
	if !reflect.DeepEqual(*headerExpect, *actual) {
		fmt.Printf("expected=%v, actual=%v\n", headerExpect, actual)
	}

	// [5],[5]
	// [1,9], [9]
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var headNode, nextNode *ListNode
	carry, remainder := 0, 0
	for l1 != nil || l2 != nil || carry > 0 {
		var p, q ListNode
		if l1 != nil {
			p = *l1
		}

		if l2 != nil {
			q = *l2
		}

		remainder = (p.Val + q.Val + carry) % 10

		newNode := ListNode{
			Val:  remainder,
			Next: nil,
		}
		if nextNode == nil {
			nextNode = &newNode
			headNode = &newNode
		} else {
			nextNode.Next = &newNode
			nextNode = nextNode.Next
		}
		carry = (p.Val + q.Val + carry) / 10

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

	}

	nextNode.Next = nil
	return headNode
}

func (e AddTwoNumbers) Print() {
	fmt.Println(`
	func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
		var headNode, nextNode *ListNode
		carry, remainder := 0, 0
		for l1 != nil || l2 != nil || carry > 0 {
			var p,q ListNode
			if l1 != nil {
				p = *l1
			}
	
			if l2 != nil {
				q = *l2
			}
	
			remainder = (p.Val + q.Val + carry) % 10
	
			newNode := ListNode{
				Val:  remainder,
				Next: nil,
			}
			if nextNode == nil {
				nextNode = &newNode
				headNode = &newNode
			} else {
				nextNode.Next = &newNode
				nextNode = nextNode.Next
			}
			carry = (p.Val + q.Val + carry) / 10
	
			if l1 != nil {
				l1 = l1.Next
			}
	
			if l2 != nil {
				l2 = l2.Next
			}
	
		}
	
		nextNode.Next = nil
		return headNode
	}
	
	`)

}
