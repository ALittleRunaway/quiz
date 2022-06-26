package game

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func printResults(total int, right int) {
	percent := (100 * right) / total

	var addMessage string
	switch {
	case percent >= 90:
		addMessage = "Amazing!"
	case (percent >= 60) && (percent < 90):
		addMessage = "Not Bad!"
	case (percent >= 30) && (percent < 60):
		addMessage = "Could be better."
	case percent < 30:
		addMessage = "Pretty bad. Better luck next time."
	}

	fmt.Printf("\n=========\n"+
		"Finished! Your score is %v of %v.\n", right, total)
	fmt.Println(addMessage)
	os.Exit(0)
}

func gameLoop(limit time.Duration, questions []Question) {

	ctx, cancel := context.WithTimeout(context.Background(), limit*time.Second)
	defer cancel()

	var correct int
	var total = len(questions)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)

		for i, questionAnswer := range questions {
			for question, answer := range questionAnswer {

				fmt.Printf("%v) %s = ", i+1, question)

				scanner.Scan()
				var userInput = scanner.Text()

				if clearString(userInput) == answer {
					correct += 1
				}
			}
		}
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			printResults(total, correct)
		}
	}
}

func StartGame() {
	cfg := initConfig()

	questions, err := readFile(cfg.quizFile)
	if err != nil {
		handleError(err.Error())
	}

	if cfg.shuffle {
		questions = shuffleQuestions(questions)
	}

	gameLoop(cfg.limit, questions)
}
