package collection

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T){
	t.Run("sum of two slices", func(t *testing.T) {
		got := SumAll([]int{1,1}, []int{1,1,1,1})
		expected := []int{2, 4}

		// Below code is not type as DeepEqual does not care about type
		//  but it will return false if types are not equal
		if !reflect.DeepEqual(got, expected){
			t.Errorf("got %v expected %v", got, expected)
		}
	})

	t.Run("sum of 5 slices", func(t *testing.T) {
		got := SumAll([]int{1,1}, []int{1,1,1,1}, []int{1,1,1}, []int{1,1,1,1,1}, []int{1})
		expected := []int{2, 4, 3, 5, 1}

		// Below code is not type as DeepEqual does not care about type
		//  but it will return false if types are not equal
		if !reflect.DeepEqual(got, expected){
			t.Errorf("got %v expected %v", got, expected)
		}
	})

}