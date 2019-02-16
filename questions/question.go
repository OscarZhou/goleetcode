package questions

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
}
