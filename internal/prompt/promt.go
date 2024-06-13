// A module used to accept user input from the cli.

package prompt

import (
	"fmt"
	"strings"
)

func UserPrompt() (reponse string) {

	var (
		user_input string
	)

	fmt.Scanf("%s", &user_input)

	return strings.TrimSpace(user_input)
}
