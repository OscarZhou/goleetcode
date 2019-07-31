package main

import (
	"flag"
	"fmt"
	"goleetcode/questions"
	"log"

	track "github.com/OscarZhou/gotrack"
)

type QuestionName string

const (
	QuestionJewelsAndStones QuestionName = "Jewels and Stones"
	QuestionValidNumber     QuestionName = "Valid Number"
	QuestionStringtoInteger QuestionName = "String to Integer"
	QuestionTwoSum          QuestionName = "Two Sum"
	QuestionAddTwoNumbers   QuestionName = "Add Two Numbers"
)

// Param is the collection of all avaiable parameters
type Param struct {
	Number     int
	NumberDesp string
	// Question Specific question on leet code
	Question string
	// OperationDesp Description for the field Question
	QuestionDesp string
	// Debug Identifier of debug mode
	Debug bool
	// DebugDesp Description for the field Debug
	DebugDesp string
	Run       bool
	RunDesp   string
	Print     bool
	PrintDesp string
	Help      bool
	HelpDesp  string
}

var (
	param = Param{
		Number:       0,
		NumberDesp:   "The question number that you want to check",
		Question:     "",
		QuestionDesp: "The question that you want to check",
		Debug:        true,
		DebugDesp:    "Display the debug information",
		Run:          false,
		RunDesp:      "Run the test case",
		Print:        false,
		PrintDesp:    "Print the sample code",
		Help:         false,
		HelpDesp:     "Print the commands",
	}
)

func main() {
	// Parse the flags
	flag.IntVar(&(param.Number), "number", param.Number, param.NumberDesp)
	flag.IntVar(&(param.Number), "n", param.Number, param.NumberDesp+" (shorthand)")
	flag.StringVar(&(param.Question), "question", param.Question, param.QuestionDesp)
	flag.StringVar(&(param.Question), "q", param.Question, param.QuestionDesp+" (shorthand)")
	flag.BoolVar(&(param.Debug), "debug", param.Debug, param.DebugDesp)
	flag.BoolVar(&(param.Debug), "dbg", param.Debug, param.DebugDesp+" (shorthand)")
	flag.BoolVar(&(param.Run), "run", param.Run, param.RunDesp)
	flag.BoolVar(&(param.Run), "r", param.Run, param.RunDesp+" (shorthand)")
	flag.BoolVar(&(param.Print), "print", param.Print, param.PrintDesp)
	flag.BoolVar(&(param.Print), "p", param.Print, param.PrintDesp+" (shorthand)")
	flag.BoolVar(&(param.Help), "help", param.Help, param.HelpDesp)
	flag.BoolVar(&(param.Help), "h", param.Help, param.HelpDesp+" (shorthand)")
	flag.Parse()

	switch flag.Arg(0) {
	case "list", "ls":
		fmt.Println("\nQuestion List:")
		fmt.Println("------------------------------------")
		questions := registerQuestions(nil)
		for k := range questions {
			questions[k].Init()
			questions[k].PrintTitle()
		}
		fmt.Println("------------------------------------")
		return
	case "version", "v":
		fmt.Println("goleetcode:0.1.1")
		return
	case "help", "h":
		fmt.Println("number, n:\t\t\t\tThe question number that you want to check")
		fmt.Println("question, q:\t\t\t\tThe question that you want to check")
		fmt.Println("debug, dbg:\t\t\t\tDisplay the debug information")
		fmt.Println("run, r:\t\t\t\t\tRun the test case")
		fmt.Println("print, p:\t\t\t\tPrint the sample code")
		fmt.Println("help, h:\t\t\t\tPrint the commands")
		return
	}

	t, err := track.New(track.Config{
		Debug:   true,
		AsynLog: false,
	})

	if err != nil {
		log.Fatal(err)
	}

	q := registerQuestions(t)
	for k := range q {
		q[k].Init()
	}

	var (
		target questions.Questioner
		ok     bool
	)
	if param.Question != "" {
		target, ok = questions.TitleMap[param.Question]
		if !ok {
			log.Fatal("Illegal question name")
		}

		if param.Print {
			target.Print()
		}
	}

	if param.Number != 0 {
		target, ok = questions.NumberMap[param.Number]
		if !ok {
			log.Fatal("Illegal question number")
		}

		if param.Print {
			target.Print()
		}
	}

	if param.Run {
		target.Run()
	}

}

func registerQuestions(t *track.Track) map[QuestionName]questions.Questioner {
	q := make(map[QuestionName]questions.Questioner)
	q[QuestionJewelsAndStones] = &questions.JewelAndStone{Track: t}
	q[QuestionValidNumber] = &questions.ValidNumber{Track: t}
	q[QuestionStringtoInteger] = &questions.StringToInteger{Track: t}
	q[QuestionTwoSum] = &questions.TwoSum{Track: t}
	q[QuestionAddTwoNumbers] = &questions.AddTwoNumbers{Track: t}

	return q
}
