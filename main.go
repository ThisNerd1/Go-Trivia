package main

import "fmt"

type Question struct{
	ques string
	answer string
}

var quiz = []Question{}

func createQuestion(question string, answer string){ 
	var x Question
	x.ques = question
	x.answer = answer
	quiz = append(quiz, x)
}

func createQuiz(){
	createQuestion("Fugu is made from what poisonous animal?", "pufferfish")
	createQuestion("What is one of the most expensive spices in the world?", "saffron")
	createQuestion("What is the biggest island in the world?", "greenland")
	createQuestion("What is the largest selling vodka brand in the world?", "smirnoff")
	createQuestion("Marzipan is made from which nuts?", "almonds")
}

func main() {
	createQuiz()
	userScore := 0
	for i := 0; i < len(quiz); i++ {
		fmt.Println(quiz[i].ques)
		var userAnswer string
		fmt.Scan(&userAnswer)
		if userAnswer == quiz[i].answer {
			fmt.Println("Correct!")
			userScore++
			} else {
				fmt.Println("Wrong!")
		}	
	}
	fmt.Println("Your score is: ", userScore, "Thanks for playing, bye!")
}
