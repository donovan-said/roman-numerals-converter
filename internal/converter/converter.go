/*
Algorithm derived from:
- https://www.noobgeek.in/blogs/roman-to-integer-leetcode-golang-solution
*/

package converter

import "fmt"

var rom_mapping = map[string]int{
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
	dec_val := 0

	// Store the previous roman symbol, so as to handle subtraction cases
	prev_num := ""

	// Loop through each symbol defined in the user_input
	for _, v := range user_input {

		if rom_mapping[string(v)] > rom_mapping[prev_num] {
			/*
				If current value is great than the previous
				value, subtract previous value from current
				value, e.g.
				IV (1, 5) = 5 - 1 = 4
			*/
			dec_val = rom_mapping[string(v)] - rom_mapping[prev_num]
		} else {
			/*
				Add the current symbols value to the result,
				e.g.
				XI (10, 1) = 10 + 1 = 11
			*/
			dec_val += rom_mapping[string(v)]
		}

		prev_num = string(v)
	}

	fmt.Println(">> Converted decimal value =", dec_val)

	return dec_val
}
