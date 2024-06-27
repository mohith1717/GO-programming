package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

type TriviaQuestion struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   string   `json:"correct_answer"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	correctAnswers := 0

	fmt.Println("Welcome to the Go Quiz Game!")

	// Fetch questions from online API
	questions, err := fetchQuestions()
	if err != nil {
		fmt.Println("Failed to fetch questions:", err)
		return
	}

	for i, question := range questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, question.Question)

		for _, option := range question.Options {
			fmt.Println(option)
		}

		fmt.Print("Your answer: ")
		userAnswer, _ := reader.ReadString('\n')
		userAnswer = strings.ToLower(strings.TrimSpace(userAnswer))

		if strings.ToLower(question.Answer) == strings.ToLower(userAnswer) {
			fmt.Println("Correct!")
			correctAnswers++
		} else {
			fmt.Printf("Incorrect! The correct answer is: %s\n", question.Answer)
		}
	}

	fmt.Printf("\nYou got %d out of %d questions correct!\n", correctAnswers, len(questions))

	if correctAnswers == len(questions) {
		fmt.Println("Congratulations! You got all the questions correct!")
	} else {
		fmt.Println("Better luck next time!")
	}
}

func fetchQuestions() ([]TriviaQuestion, error) {
	resp, err := http.Get("https://opentdb.com/api.php?amount=4&category=9&difficulty=easy&type=multiple")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Results []struct {
			Question string   `json:"question"`
			Options  []string `json:"incorrect_answers"`
			Answer   string   `json:"correct_answer"`
		} `json:"results"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	var questions []TriviaQuestion
	for _, result := range data.Results {
		options := append(result.Options, result.Answer)
		questions = append(questions, TriviaQuestion{
			Question: result.Question,
			Options:  shuffle(options), // Shuffle the options
			Answer:   result.Answer,
		})
	}

	return questions, nil
}

// Shuffle options to make the game more interesting
func shuffle(options []string) []string {
	for i := range options {
		j := rand.Intn(i + 1)
		options[i], options[j] = options[j], options[i]
	}
	return options
}
