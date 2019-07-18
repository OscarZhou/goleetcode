package questions

import (
	"fmt"
	"reflect"

	track "github.com/OscarZhou/gotrack"
)

type TwoSum struct {
	Question
	Track *track.Track
}

func (e *TwoSum) Init() {
	e.No = 1
	e.Title = "Two Sum"
	e.FullTitle = "Two Sum"
	e.URL = "https://leetcode.com/problems/two-sum/"
	e.Level = LevelEasy
	e.FuncName = "twoSum"
	NumberMap[e.No] = e
	TitleMap[e.Title] = e
}

func (e TwoSum) PrintTitle() {
	fmt.Printf("%d, %s\n", e.No, e.Title)
}

func (e TwoSum) PrintDetail() {
	fmt.Printf("%d, %s\n", e.No, e.Title)
}

func (e TwoSum) Run() error {
	e.Track.Start()
	defer e.Track.End()

	nums := []int{0, 4, 3, 0}
	target := 0
	expect := []int{0, 3}
	actual := twoSum(nums, target)
	if !reflect.DeepEqual(expect, actual) {
		fmt.Printf("expected=%v, actual=%v\n", expect, actual)
	}
	nums = []int{-1, -2, -3, -4, -5}
	target = -8
	expect = []int{2, 4}
	actual = twoSum(nums, target)
	if !reflect.DeepEqual(expect, actual) {
		fmt.Printf("expected=%v, actual=%v\n", expect, actual)
	}

	return nil
}

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

func (e TwoSum) Print() {
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
