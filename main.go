package main

import (
	"fmt"
)

type Quiz struct {
    Question string
    Answer   string
}
type CsvFile struct {
	quizs []Quiz
}


var f = CsvFile{
	quizs: []Quiz{
		{Question: "5+5", Answer: "10"},
		{Question: "5+6", Answer: "11"},
		{Question: "5+7", Answer: "12"},
		{Question: "5+8", Answer: "13"},
	},
}


func main() {
	
	var userInpt string
	for idx, q := range f.quizs {
		fmt.Printf("Question number %d: how much is %s ?\n", idx, q.Question)
		fmt.Scanln(&userInpt)
		messg := ""
		if userInpt == q.Answer {
			messg = "Correct answer!"
		} else {
			messg = "Wrong..."
		}
		fmt.Printf("User inpuut is: %s, where as the correct answer is: %s. %s\n", userInpt, q.Answer, messg)
	}
	
}