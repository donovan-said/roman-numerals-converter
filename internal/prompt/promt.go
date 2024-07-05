// A module used to accept user input from the cli.

package prompt

import (
	"fmt"
	"strings"
)

// UserPrompt is used to request user input.
func UserPrompt() string {
	var (
		userInput string
	)

	_, err := fmt.Scanf("%s", &userInput)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(userInput)
}
