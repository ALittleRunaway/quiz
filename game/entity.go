package game

import "time"

type Config struct {
	quizFile string
	limit    time.Duration
	shuffle  bool
}

type Question map[string]string
