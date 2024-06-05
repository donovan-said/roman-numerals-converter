/*

The purpose of this tool is convert Roman numerals to decimal numerals.


Below is a mapping of Roman numerals to decimal numerals:

I - 1
V - 5
X - 10
L - 50
C - 100
D - 500
M - 1000

This list has been obtained from https://en.wikipedia.org/wiki/Roman_numerals

*/

package main

import (
	"os"

	"github.com/donovan-said/roman-numerals-converter/internal/converter"
	"github.com/donovan-said/roman-numerals-converter/internal/prompt"
)

func main() {
	input, val := prompt.UserPrompt()
	if !val {
		os.Exit(1)
	} else {
		converter.Converter(input)
	}

}
