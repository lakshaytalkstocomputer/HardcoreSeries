package roman

import "strings"

// Write a function that converts Arabic Number (0 to 9) to a Roman numeral.
//   I is "one"          , II is "two"  , III is "three".
//   V is "five"         , Iv is "four" , X is "ten"
//   MCMLXXXIV is "1984" .
//
// Key skill for software developers is to try and identify "thin vertical slices"
//   and then iterating. the TDD workflow helps facilitate iterative development.

// This function converts Arabic Number to Roman Number
//   Romans believed in DRY too,
//   They thought repeating characters would become hard to read and
//   count. So rule of the thumb with Roman numeral is:
//   You can't have same character repeated more than 3 times in a row.
//   Instead you take the next highest symbol and then "subtract" by putting a
//   symbol to left of it.
//   Not all symbols can be used as subtractors only 1 (I), X(10), C(100)
func ConvertToRoman(number int) string{
	var output strings.Builder

	for number > 0 {
		switch {
		case number > 9:
			output.WriteString("X")
			number -= 10
		case number > 8:
			output.WriteString("IX")
			number -= 9
		case number > 4:
			output.WriteString("V")
			number -= 5
		case number > 3:
			output.WriteString("IV")
			number -= 4
		default:
			output.WriteString("I")
			number--
		}
	}
	return output.String()
}


