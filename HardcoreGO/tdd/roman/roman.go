package roman

import "strings"

// Write a function that converts Arabic Number (0 to 9) to a Roman numeral.
//   I is "one"          , II is "two"  , III is "three".
//   V is "five"         , Iv is "four" , X is "ten"
//   MCMLXXXIV is "1984" .
//
// Key skill for software developers is to try and identify "thin vertical slices"
//   and then iterating. the TDD workflow helps facilitate iterative development.

type RomanNumeral struct{
	Value int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}


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

	for _, numeral := range allRomanNumerals{
		for number >= numeral.Value{
			output.WriteString(numeral.Symbol)
			number -= numeral.Value
		}
	}

	return output.String()
}


