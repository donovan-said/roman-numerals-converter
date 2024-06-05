/*

A module used to accept user input and validate that the input conforms to the
Rooman numerals.

*/

package prompt

import (
	"fmt"
	"slices"
	"strings"
)

var roman_numerals = []string{"I", "V", "X", "L", "C", "D", "M"}

func Validator(roman_list []string, user_list string) bool {
	// Exmaple: https://gosamples.dev/slice-contains/

	for _, s := range user_list {
		/*
		 - Convert rune to string
		 - Convert string to uppercase
		 - Ensureing string is in list of Roman numerals
		*/
		val := slices.Contains(roman_list, strings.ToUpper(string(s)))

		if !val {
			return false
		}
	}

	return true
}

func UserPrompt() (reponse string, validator bool) {

	// Obtaining user input via the cli
	var (
		user_input string
	)
	fmt.Println("Enter Roman Numeral from list of I, V, X, L, C, D, M: ")
	fmt.Scanf("%s", &user_input)

	/*
	Split user input into slice so as to compare each element against the
	approved list of Roman numerals,
	*/

	val := Validator(roman_numerals, user_input)

	if !val {
		fmt.Printf("%s is not a valid Roman numeral!", user_input)
	} else {
		fmt.Println("You entered Roman numeral:", user_input)
	}

	return strings.TrimSpace(user_input), val
}
