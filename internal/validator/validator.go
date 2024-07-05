/*

A module used validate if a provided string is available in a provided array of
strings.

*/

package validator

import (
	"fmt"
	"slices"
)

func Validator(validationList []string, userInput string) error {
	val := slices.Contains(validationList, userInput)

	if !val {
		return fmt.Errorf("%s is an invalid value", userInput)
	}

	return nil
}
