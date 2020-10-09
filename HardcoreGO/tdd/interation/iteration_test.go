package iteration

import (
	"fmt"
	"testing"
)

func TestRepeater(t *testing.T){
	checkerFunc := func(t *testing.T, expected string, got string) {
		t.Helper()

		if got != expected{
			t.Errorf("Expected %q but got %q", expected, got)
		}
	}

	t.Run("Return a 6 times", func(t *testing.T) {
		got 		:= Repeat("a", 6)
		expected 	:= "aaaaaa"
		checkerFunc(t, expected, got)
	})

	t.Run("Return b 6 times", func(t *testing.T) {
		got 		:= Repeat("b", 6)
		expected 	:= "bbbbbb"
		checkerFunc(t, expected, got)
	})

	t.Run("Return c 8 times", func(t *testing.T) {
		got := Repeat("c", 9)
		expected := "ccccccccc"
		checkerFunc(t, expected, got)
	})

	t.Run("Return d 0 times", func(t *testing.T) {
		got := Repeat("d", 0)
		expected := ""
		checkerFunc(t, expected, got)
	})

	t.Run("Return e 0 times when negative number provided", func(t *testing.T) {
		got := Repeat("e", -10)
		expected := ""
		checkerFunc(t, expected, got)
	})


}

// This function will benchmark the function
// * Use go test -bench=. to run benchmarking
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i<b.N; i++{
		Repeat("a", 6)
	}
}

// This function will add example for the function
func ExampleRepeat() {
	repeatedText := Repeat("a", 10)
	fmt.Println("Output :", repeatedText)
}