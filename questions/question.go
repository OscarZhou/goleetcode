package questions

import (
	"errors"
	"fmt"
)

type LevelType string

const (
	LevelEasy   LevelType = "Easy"
	LevelMedium LevelType = "Medium"
	LevelHard   LevelType = "Hard"
)

type Questioner interface {
	Init()
	Print()
	PrintTitle()
	PrintDetail()
	Run() error
	// Output()
	// Log()
}

// Question records the basic information of the question
type Question struct {
	No        int
	Title     string
	FullTitle string
	URL       string
	Level     LevelType
	FuncName  string
}

func (Question) Init() {
}

func (Question) Print() {
	fmt.Println("no contents")
}

func (Question) PrintTitle() {
	fmt.Println("no contents")
}

func (Question) PrintDetail() {
	fmt.Println("no contents")
}

func (Question) Run() error {
	return errors.New("method is not allowed")
}

// func (Question) Output() error {
// 	return errors.New("method is not allowed")
// }

// func (Question) Log() error {
// 	return errors.New("method is not allowed")
// }
