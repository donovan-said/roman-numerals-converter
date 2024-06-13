/*

A module used validate if a provided string is available in a provided array of
strings.

*/

package validator

import (
	"fmt"
	"slices"
)

func Validator(validation_list []string, user_input string) error {

	val := slices.Contains(validation_list, user_input)

	if !val {
		return fmt.Errorf("%s is an invalid value", user_input)
	}

	return nil
}
