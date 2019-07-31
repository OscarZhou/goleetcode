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

	l1 := &ListNode{Val: 3, Next: nil}
	header1 := l1
	l1.Next = &ListNode{Val: 4, Next: nil}
	l1 = l1.Next
	l1.Next = &ListNode{Val: 2, Next: nil}
	l1 = l1.Next
	l1.Next = nil

	l2 := &ListNode{Val: 4, Next: nil}
	header2 := l2
	l2.Next = &ListNode{Val: 6, Next: nil}
	l2 = l2.Next
	l2.Next = &ListNode{Val: 5, Next: nil}
	l2 = l2.Next
	l2.Next = nil

	expect := &ListNode{Val: 8, Next: nil}
	headerExpect := expect
	expect.Next = &ListNode{Val: 0, Next: nil}
	expect = expect.Next
	expect.Next = &ListNode{Val: 7, Next: nil}
	expect = expect.Next
	expect.Next = nil

	actual := addTwoNumbers(header1, header2)
	if !reflect.DeepEqual(*headerExpect, *actual) {
		fmt.Printf("expected=%v, actual=%v\n", headerExpect, actual)
	}

	l1 = &ListNode{Val: 1, Next: nil}
	header1 = l1
	l1.Next = &ListNode{Val: 8, Next: nil}
	l1 = l1.Next
	l1.Next = nil

	l2 = &ListNode{Val: 0, Next: nil}
	l2.Next = nil

	header2 = l2
	expect = &ListNode{Val: 1, Next: nil}
	headerExpect = expect
	expect.Next = &ListNode{Val: 8, Next: nil}
	expect = expect.Next
	expect.Next = nil

	actual = addTwoNumbers(header1, header2)
	if !reflect.DeepEqual(*headerExpect, *actual) {
		fmt.Printf("expected=%v, actual=%v\n", headerExpect, actual)
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	nextNode := new(ListNode)
	headNode := new(ListNode)
	headNode = nextNode
	carry, remainder := 0, 0
	p := ListNode{Val: 0, Next: nil}
	q := ListNode{Val: 0, Next: nil}
	for l1 != nil || l2 != nil {
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
		}
		carry = (p.Val + q.Val + carry) / 10

		if l1.Next != nil {
			l1 = l1.Next
		}

		if l2.Next != nil {
			l2 = l2.Next
		}
	}

	nextNode.Next = nil
	return headNode
}

func (e AddTwoNumbers) Print() {
	fmt.Println(`
	func twoSum(nums []int, target int) []int {
		length := len(nums)
		for i := 0; i < length; i++ {
			for j := i + 1; j < length; j++ {
				if nums[i]+nums[j] == target {
					return []int{i, j}
				}
			}
		}
	
		return []int{}
	}
	`)

}
