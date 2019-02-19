package questions

import "errors"

type Questioner interface {
	Print()
	Output()
	Log()
}

// Question records the basic information of the question
type Question struct {
	No          int
	Name        string
	Description string
	URL         string
}

func (Question) Print() error {
	return errors.New("method is not allowed")
}

func (Question) Output() error {
	return errors.New("method is not allowed")
}

func (Question) Log() error {
	return errors.New("method is not allowed")
}
