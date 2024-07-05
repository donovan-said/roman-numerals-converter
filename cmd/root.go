// A module used to request user input.

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/donovan-said/roman-numerals-converter/internal/converter"
	"github.com/donovan-said/roman-numerals-converter/internal/prompt"
	"github.com/donovan-said/roman-numerals-converter/internal/validator"
)

var roman_numerals = []string{"I", "V", "X", "L", "C", "D", "M"}

func Execute() {

	fmt.Println(">> Enter Roman Numeral from list of [I, V, X, L, C, D, M]: ")
	input := prompt.UserPrompt()
	fmtd_input := strings.ToUpper(input)
	fmt.Printf(">> Validating input %s...\n", fmtd_input)

	/*
		Validating each element of the user input, ensuring that it is a valid
		Roman numeral.
	*/
	for _, s := range fmtd_input {
		// Convert rune to string
		err := validator.Validator(roman_numerals, string(s))
		if err != nil {
			fmt.Println(">> Error encountered:", err)
			os.Exit(1)
		}
	}

	converter.Converter(fmtd_input)
}
