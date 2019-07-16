package questions

import (
	"fmt"
	"strings"

	track "github.com/OscarZhou/gotrack"
)

type StringToInteger struct {
	Question
	Track *track.Track
}

func (e *StringToInteger) Init() {
	e.No = 8
	e.Title = "String to Integer"
	e.FullTitle = "String to Integer"
	e.URL = "https://leetcode.com/problems/string-to-integer-atoi/"
	e.Level = LevelMedium
	e.FuncName = "myAtoi"
	NumberMap[e.No] = e
	TitleMap[e.Title] = e
}

func (e StringToInteger) Print() {
	fmt.Println(`
	func numJewelsInStones(J string, S string) int {
		count := 0
		for _, s := range []rune(S) {
			for _, j := range []rune(J) {
				if s == j {
					count++
					break
				}
			}
		}
		return count
	}
	`)

}

func (e StringToInteger) PrintTitle() {
	fmt.Printf("%d, %s\n", e.No, e.Title)
}

func (e StringToInteger) PrintDetail() {
	fmt.Printf("%d, %s\n", e.No, e.Title)
}

func (e StringToInteger) Run() error {
	e.Track.Start()
	defer e.Track.End()

	sArray := []string{
		"42",
		"    -42",
		"  42",
		"4193 with word",
		"-91283472332",
		"",
		"+",
		"-000000000000001",
		"+-2",
		"010",
		"9223372036854775808",
		"2147483648",
		"-2147483647",
		"-912834723328743294",
		"0-1",
		"    0000000000000   ",
		"-5-",
		"18446744073709551617",
		"  0000000000012345678",
	}

	expectedArray := []int{
		42,
		-42,
		42,
		4193,
		-2147483648,
		0,
		0,
		-1,
		0,
		10,
		2147483647,
		2147483647,
		-2147483647,
		-2147483648,
		0,
		0,
		-5,
		2147483647,
		12345678,
	}

	for i, v := range sArray {
		actual := myAtoi(v)
		if expectedArray[i] != actual {
			fmt.Printf("index=%d, expected=%v, actual=%v\n", i, expectedArray[i], actual)
		}

	}
	return nil
}

func myAtoi(str string) int {
	s := strings.TrimLeft(str, " ")
	i, state, initalNumber, positive := 0, 0, true, true
	count := 0
	var result int64 = 0

	var max int64 = 2147483647  // int(math.Pow(2, 31)) - 1
	var min int64 = -2147483648 // int(math.Pow(2, 31) * (-1))

	castChars := make(map[rune]int64)
	castChars['0'] = 0
	castChars['1'] = 1
	castChars['2'] = 2
	castChars['3'] = 3
	castChars['4'] = 4
	castChars['5'] = 5
	castChars['6'] = 6
	castChars['7'] = 7
	castChars['8'] = 8
	castChars['9'] = 9
	castChars['+'] = 1
	castChars['-'] = -1

	chars := []rune(s)
	if len(chars) == 0 {
		return 0
	}
	for i < len(chars) {
		switch chars[i] {
		case '0':
			if !initalNumber {
				result = result*10 + castChars[chars[i]]
				count++
			}
			state = 2
			i++
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if state == 0 || state == 1 || state == 3 || state == 2 {
				state = 2
				if initalNumber {
					result = 1
					result = result * castChars[chars[i]]
					initalNumber = false
				} else {
					result = result*10 + castChars[chars[i]]
				}
				count++
				i++
			} else {
				return 0
			}

		case '+', '-':
			if state == 0 {
				state = 3
				if castChars[chars[i]] < 0 {
					positive = false
				}
				i++
			} else if state == 2 {
				i = len(chars)
			} else {
				return 0
			}
		default:
			if i == 0 {
				return 0
			} else {
				i = len(chars)
			}
		}
		if count > 10 {
			break
		}
	}
	if state == 2 {
		if result > max || result < min {
			if positive {
				return int(max)
			} else {
				return int(min)
			}
		}

		if positive {
			return int(result)
		} else {
			return int(result) * (-1)
		}
	}
	return 0
}
