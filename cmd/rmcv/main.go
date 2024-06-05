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
	"github.com/donovan-said/roman-numerals-converter/internal/prompt"
)

func main() {
	//----------------------------------------------------------------------
	// User input is used to decide whether to populate the database or not
	prompt.UserPrompt()
}
