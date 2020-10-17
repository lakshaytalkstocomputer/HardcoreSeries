package roman

import (
	"fmt"
	"testing"
	"testing/quick"
)

// Use the uneasy feeling to write a new test to force us to write slightly less dumb code.

// Create table test with the test cases
var tt = []struct{
			Description string
			Arabic 		uint16
			Expected 	string
			}{
				{"1 to I", 1, "I"},
				{"2 to II", 2, "II"},
				{"3 to III", 3, "III"},
				{"4 to IV", 4, "IV"},
				{"5 to V", 5, "V"},
				{"6 to VII", 6, "VI"},
				{"7 to VII", 7, "VII"},
				{"8 to VIII", 8, "VIII"},
				{"9 to IX", 9, "IX"},
				{"10 to X", 10, "X"},
				{"14 to XIV", 14, "XIV"},
				{"18 to XVIII", 18, "XVIII"},
				{"20 to XX", 20, "XX"},
				{"39 to XXXIX", 39, "XXXIX"},
				{"40 to XL", 40, "XL"},
				{"47 to XLVII", 47, "XLVII"},
				{"49 to XLIX", 49, "XLIX"},
				{"50 to L", 50, "L"},
				{"1984 to MCMLXXIV", 1984, "MCMLXXXIV"},
			}

func TestRomanNumber(t *testing.T){

	for _, testData := range tt{
		t.Run(testData.Description, func(t *testing.T) {

			var got	string	 = ConvertToRoman(uint16(testData.Arabic))

			if got != testData.Expected{
				t.Errorf("got %v, expected %v", got, testData.Expected)
			}
		})
	}

}

func TestRomanToArabic(t *testing.T){
	for _, test := range tt{
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Expected, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Expected)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount:1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
