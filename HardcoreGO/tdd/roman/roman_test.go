package roman

import "testing"

// Use the uneasy feeling to write a new test to force us to write slightly less dumb code.


func TestRomanNumber(t *testing.T){

	// Create table test with the test cases
	tt := []struct{
		Description string
		Arabic 		int
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
	}

	for _, testData := range tt{
		t.Run(testData.Description, func(t *testing.T) {

			var got	string	 = ConvertToRoman(testData.Arabic)

			if got != testData.Expected{
				t.Errorf("got %v, expected %v", got, testData.Expected)
			}
		})
	}

}