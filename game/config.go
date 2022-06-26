package game

import (
	"github.com/spf13/viper"
)

func initConfig() Config {
	conf := Config{
		quizFile: viper.GetString("quiz_file"),
		limit:    viper.GetDuration("limit"),
		shuffle:  viper.GetBool("shuffle"),
	}
	return conf
}
