package iteration

import "testing"

func TestRepeater(t *testing.T){
	checkerFunc := func(t *testing.T, expected string, got string) {
		t.Helper()

		if got != expected{
			t.Errorf("Expected %q but got %q", expected, got)
		}
	}

	t.Run("Return a 6 times", func(t *testing.T) {
		got 		:= Repeat("a")
		expected 	:= "aaaaaa"
		checkerFunc(t, expected, got)
	})

	t.Run("Return b 6 times", func(t *testing.T) {
		got 		:= Repeat("b")
		expected 	:= "bbbbbb"
		checkerFunc(t, expected, got)
	})
}