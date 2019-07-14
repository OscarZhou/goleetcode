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
}

var (
	param = Param{
		Number:       0,
		NumberDesp:   "The question number that you want to check",
		Question:     "",
		QuestionDesp: "The question that you want to check",
		Debug:        true,
		DebugDesp:    "Display the debug information",
		Run:          true,
		RunDesp:      "Run the test case",
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
	flag.BoolVar(&(param.Run), "run", param.Debug, param.RunDesp)
	flag.BoolVar(&(param.Run), "r", param.Debug, param.RunDesp+" (shorthand)")
	flag.Parse()

	switch flag.Arg(0) {
	case "list", "ls":
		fmt.Println("\nQuestion List:")
		fmt.Println("------------------------------------")
		questions := registerQuestions(nil)
		for k := range questions {
			questions[k].Init()
			questions[k].PrintTitle()
			// ptr := reflect.New(reflect.TypeOf(questions[k]).Elem())
			// methodInit := ptr.MethodByName("Init")
			// if !methodInit.IsValid() {
			// 	info := fmt.Sprintf("%s fails to Init", k)
			// 	log.Fatal(info)
			// }
			// args := make([]reflect.Value, 0)
			// methodInit.Call(args)
			// methodPrintTitle := ptr.MethodByName("PrintTitle")
			// if !methodInit.IsValid() {
			// 	info := fmt.Sprintf("%s fails to PrintTitle", k)
			// 	log.Fatal(info)
			// }
			// methodPrintTitle.Call(args)

		}
		fmt.Println("------------------------------------")
		return
	case "version", "v":
		fmt.Println("goleetcode:0.1.1")
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

		target.Print()
	}

	if param.Number != 0 {
		fmt.Println(questions.NumberMap)
		target, ok = questions.NumberMap[param.Number]
		if !ok {
			log.Fatal("Illegal question number")
		}

		target.Print()
	}

	if param.Run {
		target.Run()
	}

}

func registerQuestions(t *track.Track) map[QuestionName]questions.Questioner {
	q := make(map[QuestionName]questions.Questioner)
	q[QuestionJewelsAndStones] = &questions.JewelAndStone{Track: t}
	q[QuestionValidNumber] = &questions.ValidNumber{Track: t}

	return q
}
