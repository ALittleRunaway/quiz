package cmd

/*
Copyright © 2022 Maria Petrova marycool674@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"quiz/game"
)

var (
	quizFile string
	limit    int
	shuffle  bool
)

var rootCmd = &cobra.Command{
	Use:     "quiz",
	Short:   "A CLI quiz game",
	Example: "quiz -s -l 10 -f questions.csv",
	Long: `Make the .csv file with two columns: 1) questions, 2) answers.
The app will parse that and give you the opportunity to answer those questions.
`,

	Run: func(cmd *cobra.Command, args []string) {
		game.StartGame()
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVarP(&quizFile, "game-file", "f", "problems.csv",
		"Sets the path to file with questions in .csv format& default: problems.cev")
	rootCmd.Flags().IntVarP(&limit, "limit", "l", 30,
		"Sets the timer limit in seconds. Default: 30 sec")
	rootCmd.Flags().BoolVarP(&shuffle, "shuffle", "s", false,
		"Shuffles the game questions")
}

func initConfig() {
	viper.Set("quiz_file", quizFile)
	viper.Set("limit", limit)
	viper.Set("shuffle", shuffle)
}
