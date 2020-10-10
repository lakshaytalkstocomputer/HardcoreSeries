package dictionary

import "testing"

func TestSearch(t *testing.T){

	t.Run("Search for known Word", func(t *testing.T) {
		dict := Dictionary{"hello":"a word to greet others"}

		got, err := dict.Search("hello")
		expected := "a word to greet others"

		if err != nil{
			t.Errorf("error should be nil but got %s",err)
		}

		if got != expected{
			t.Errorf("Got %s Expected %s", got, expected)
		}
	})

	t.Run("Searhc for Unknown Word", func(t *testing.T) {
		dict := Dictionary{}

		got, err := dict.Search("hello")

		if err == nil{
			t.Errorf("Error should not be nil")
		}

		if got != ""{
			t.Errorf("Value must be Zero Value of string")
		}

	})
}

func TestAdd(t *testing.T){
	t.Run("Testing Addition of key", func(t *testing.T) {

		// Create A Dictionary
		dict := Dictionary{}

		// Add word and definition into dictionary
		_ = dict.Add("hello", "greeting word to say hi")

		got, err := dict.Search("hello")
		expected := "greeting word to say hi"

		if err != nil {
			t.Error("Error should be nil but got :",err)
		}

		if got != expected{
			t.Errorf("Got %q expected %q",got, expected)
		}
	})

	t.Run("Addition of already present key", func(t *testing.T) {

		// Create a dicitonary
		dict := Dictionary{"hello": "greeting word"}

		err := dict.Add("hello", "some greet word")

		if err == nil {
			t.Error("Error should not be nil")
		}

	})
}