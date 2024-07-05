/*
Algorithm derived from:
- https://www.noobgeek.in/blogs/roman-to-integer-leetcode-golang-solution
*/

package converter

import "fmt"

// Converter is used to convert roman numerals to decimal numbers.
func Converter(userInput string) int {
	var romMapping = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	// Initialize the results value
	decVal := 0

	// Store the previous roman symbol, so as to handle subtraction cases
	prevNum := ""

	// Loop through each symbol defined in the userInput
	for _, v := range userInput {
		if romMapping[string(v)] > romMapping[prevNum] {
			/*
				If current value is great than the previous
				value, subtract previous value from current
				value, e.g.
				IV (1, 5) = 5 - 1 = 4
			*/
			decVal = romMapping[string(v)] - romMapping[prevNum]
		} else {
			/*
				Add the current symbols value to the result,
				e.g.
				XI (10, 1) = 10 + 1 = 11
			*/
			decVal += romMapping[string(v)]
		}

		prevNum = string(v)
	}

	fmt.Println(">> Converted decimal value =", decVal)

	return decVal
}
