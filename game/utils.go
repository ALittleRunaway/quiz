package game

import (
	"encoding/csv"
	"errors"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func clearString(line string) string {
	return strings.ToLower(strings.TrimSpace(line))
}

func readFile(filename string) ([]Question, error) {
	var questions []Question

	file, err := os.Open(filename)
	if err != nil {
		return questions, err
	}

	if filepath.Ext(filename) != ".csv" {
		handleError("file extension is not .csv")
	}

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return questions, err
		}
		if len(record) != 2 {
			return questions, errors.New("there has to be two columns in the file")
		}
		questions = append(questions, Question{clearString(record[0]): clearString(record[1])})
	}
	return questions, nil
}

func shuffleQuestions(questions []Question) []Question {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
	return questions
}
