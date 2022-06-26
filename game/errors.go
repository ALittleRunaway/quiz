package game

import (
	"fmt"
	"os"
)

func handleError(errorMessage string) {
	fmt.Println(errorMessage)
	os.Exit(1)
}
