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
}

var (
	param = Param{
		Number:       0,
		NumberDesp:   "The question number that you want to check",
		Question:     "",
		QuestionDesp: "The question that you want to check",
		Debug:        true,
		DebugDesp:    "Display the debug information",
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
	flag.Parse()

	switch flag.Arg(0) {
	case "list", "ls":
		fmt.Println("\nQuestion List:")
		fmt.Println("------------------------------------")
		questions := registerQuestions(nil)
		for k := range questions {
			fmt.Println(k)
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

	questions := registerQuestions(t)
	if param.Question != "" {
		target, ok := questions[QuestionName(param.Question)]
		if !ok {
			log.Fatal("Illegal question name")
		}

		target.Print()
	}

	if param.Number != 0 {

	}
}

func registerQuestions(t *track.Track) map[QuestionName]questions.Questioner {
	q := make(map[QuestionName]questions.Questioner)
	q[QuestionJewelsAndStones] = &questions.JewelAndStone{Track: t}
	return q
}
