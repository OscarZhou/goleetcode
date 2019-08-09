package questions

import (
	"fmt"
	"reflect"
	"strings"

	track "github.com/OscarZhou/gotrack"
)

// LSWRC means LongestSubstringWithoutRepeatingCharacters
type LSWRC struct {
	Question
	Track *track.Track
}

func (e *LSWRC) Init() {
	e.No = 3
	e.Title = "Longest Substring Without Repeating Characters"
	e.FullTitle = "Longest Substring Without Repeating Characters"
	e.URL = "https://leetcode.com/problems/longest-substring-without-repeating-characters/"
	e.Level = LevelMedium
	e.FuncName = "lengthOfLongestSubstring"
	NumberMap[e.No] = e
	TitleMap[e.Title] = e
}

func (e LSWRC) PrintTitle() {
	fmt.Printf("%d, %s\n", e.No, e.Title)
}

func (e LSWRC) PrintDetail() {
	fmt.Printf("%d, %s\n", e.No, e.Title)
}

func (e LSWRC) Run() error {
	e.Track.Start()
	defer e.Track.End()

	expect := 2
	actual := lengthOfLongestSubstring("au")
	if !reflect.DeepEqual(expect, actual) {
		fmt.Printf("expected=%v, actual=%v\n", expect, actual)
	}

	return nil
}

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length == 0 {
		return length
	}

	max, count := 1, 1
	i, j := 0, 1
	for i < length && j < length {
		k := strings.Index(s[i:j], string(s[j]))
		if k >= 0 {
			if max < count {
				max = count
			}
			i, j = i+k+1, i+k+2
			count = 1

		} else {
			count++
			j++
		}
	}
	if max < count {
		max = count
		count = 1
	}
	return max
}

func (e LSWRC) Print() {
	fmt.Println(`
	func lengthOfLongestSubstring(s string) int {
		length := len(s)
		if length == 0 {
			return length
		}
	
		max, count := 1, 1
		i, j := 0, 1
		for i < length && j < length {
			k := strings.Index(s[i:j], string(s[j]))
			if k >= 0 {
				if max < count {
					max = count
				}
				i, j = i+k+1, i+k+2
				count = 1
	
			} else {
				count++
				j++
			}
		}
		if max < count {
			max = count
			count = 1
		}
		return max
	}
	`)

}
