package questions

import (
	"fmt"
	"strings"

	track "github.com/OscarZhou/gotrack"
)

type ValidNumber struct {
	Question
	Track *track.Track
}

func (e *ValidNumber) Init() {
	e.No = 65
	e.Title = "Valid Number"
	e.FullTitle = "Valid Number"
	e.URL = "https://leetcode.com/problems/jewels-and-stones"
	e.Level = LevelHard
	e.FuncName = "validNumber"
	NumberMap[e.No] = e
	TitleMap[e.Title] = e
}

func (e ValidNumber) Print() {
	fmt.Println(`
	func isNumber(s string) bool {
		state := 0
		s = strings.TrimSpace(s)
		chars := []rune(s)
		i := 0
	
		for i < len(chars) {
			switch chars[i] {
			case '+', '-':
				if state == 0 {
					state = 1
					i++
				} else if state == 5 {
					state = 7
					i++
				} else {
					return false
				}
			case 'e':
				if state == 2 || state == 3 || state == 4 || state == 9 {
					state = 5
					i++
				} else {
					return false
				}
			case '.':
				if state == 0 || state == 1 {
					state = 8
					i++
				} else if state == 2 {
					state = 3
					i++
				} else {
					return false
				}
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				switch state {
				case 0, 1, 2:
					state = 2
					i++
				case 3, 4:
					state = 4
					i++
				case 5, 6, 7:
					state = 6
					i++
				case 8, 9:
					state = 9
					i++
				default:
					return false
				}
			default:
				return false
			}
		}
	
		if state == 2 || state == 3 || state == 4 || state == 6 || state == 9 {
			return true
		}
		return false
	}
	`)

}

func (e ValidNumber) PrintTitle() {
	fmt.Printf("%d, %s\n", e.No, e.Title)
}

func (e ValidNumber) PrintDetail() {
	fmt.Printf("%d, %s\n", e.No, e.Title)
}

func (e ValidNumber) Run() error {
	e.Track.Start()
	defer e.Track.End()

	sArray := []string{
		"0",
		" 0.1",
		"abc",
		"1 a",
		"2e10",
		" -90e3",
		" 1e",
		"e3",
		" 6e-1",
		" 99e2.5 ",
		"53.5e93",
		" --6",
		"-+3",
		"95a54e53",
		"1.",
		".3",
		"40.81",
		"46.e3",
	}

	expectedArray := []bool{
		true,
		true,
		false,
		false,
		true,
		true,
		false,
		false,
		true,
		false,
		true,
		false,
		false,
		false,
		true,
		true,
		true,
		true,
	}

	for i, v := range sArray {

		actual := validNumber(v)
		if expectedArray[i] == actual {
			fmt.Printf("index=%d, expected=%v, actual=%v\n", i, expectedArray[i], actual)
		}

	}

	return nil
}

func validNumber(s string) bool {
	return isNumber(s)
}

func isNumber(s string) bool {
	state := 0
	s = strings.TrimSpace(s)
	chars := []rune(s)
	i := 0

	for i < len(chars) {
		switch chars[i] {
		case '+', '-':
			if state == 0 {
				state = 1
				i++
			} else if state == 5 {
				state = 7
				i++
			} else {
				return false
			}
		case 'e':
			if state == 2 || state == 3 || state == 4 || state == 9 {
				state = 5
				i++
			} else {
				return false
			}
		case '.':
			if state == 0 || state == 1 {
				state = 8
				i++
			} else if state == 2 {
				state = 3
				i++
			} else {
				return false
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			switch state {
			case 0, 1, 2:
				state = 2
				i++
			case 3, 4:
				state = 4
				i++
			case 5, 6, 7:
				state = 6
				i++
			case 8, 9:
				state = 9
				i++
			default:
				return false
			}
		default:
			return false
		}
	}

	if state == 2 || state == 3 || state == 4 || state == 6 || state == 9 {
		return true
	}
	return false
}
