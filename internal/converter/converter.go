/*
Algorithm derived from:
- https://www.noobgeek.in/blogs/roman-to-integer-leetcode-golang-solution
*/

package converter

import "fmt"

var symbol_mapping = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func Converter(user_input string) (reponse int) {
	// Initialize the results value
	decimal_value := 0

	// Store the previous roman symbol, so as to handle subtraction cases
	previous_numeral := ""

	// Loop through each symbol defined in the user_input
	for _, v := range user_input {
		/*
			Checking to see if the current value is greater than the
			previous symbol.
		*/
		if symbol_mapping[string(v)] > symbol_mapping[previous_numeral] {
			// Subtract the previous value from the result
			decimal_value = decimal_value - symbol_mapping[previous_numeral]
			// Calculate the difference
			decimal_value = decimal_value + (symbol_mapping[string(v)] - symbol_mapping[previous_numeral])
		} else {
			// Add the current symbols value to the result
			decimal_value += symbol_mapping[string(v)]
		}

		previous_numeral = string(v)
	}

	fmt.Println(">> Converted decimal value =", decimal_value)

	return decimal_value
}
