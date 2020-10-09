package collection

import "testing"

func TestArraySum(t *testing.T){
	t.Run("sum of array with 5 elements", func(t *testing.T) {
		numbers := []int{1,1,1,1,1}
		got 		:= ArraySum(numbers)
		expected 	:= 5

		if expected != got{
			t.Errorf("Expected %d but got %d, give %v", expected, got, numbers)
		}
	})

	t.Run("sum of array with 10 elements", func(t *testing.T) {
		numbers := []int{1,1,1,1,1,1,1,1,1,1}
		got 		:= ArraySum(numbers)
		expected 	:= 10

		if expected != got{
			t.Errorf("Expected %d but got %d, give %v", expected, got, numbers)
		}
	})

	t.Run("sum of array with 0 elements", func(t *testing.T) {
		var numbers []int
		got 		:= ArraySum(numbers)
		expected 	:= 0

		if expected != got{
			t.Errorf("Expected %d but got %d, give %v", expected, got, numbers)
		}
	})
}
