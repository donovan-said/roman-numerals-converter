// A module used to request user input.

package cmd

import (
	"fmt"
	"os"
	"strings"

	"roman-numerals-converter/internal/converter"
	"roman-numerals-converter/internal/prompt"
	"roman-numerals-converter/internal/validator"
)

// Execute the roman conversion functionality.
func Execute() {
	var romanNumerals = []string{"I", "V", "X", "L", "C", "D", "M"}

	fmt.Println(">> Enter Roman Numeral from list of [I, V, X, L, C, D, M]: ")
	input := prompt.UserPrompt()
	fmtdInput := strings.ToUpper(input)
	fmt.Printf(">> Validating input %s...\n", fmtdInput)

	/*
		Validating each element of the user input, ensuring that it is a valid
		Roman numeral.
	*/
	for _, s := range fmtdInput {
		// Convert rune to string
		err := validator.Validator(romanNumerals, string(s))
		if err != nil {
			fmt.Println(">> Error encountered:", err)
			os.Exit(1)
		}
	}

	converter.Converter(fmtdInput)
}
