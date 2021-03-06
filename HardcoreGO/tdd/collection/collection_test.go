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

func TestSumAllTail(t *testing.T){

	// This function adds Type safety
	//   that is ignored by reflect.DeepEqual() in want parameter
	checkSums := func(t *testing.T, got, want []int){
		t.Helper()

		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum of tails of two slices", func(t *testing.T) {
		got 		:= SumAllTail([]int{1,2,2}, []int{2,3,3})
		expected	:= []int{4,6}

		checkSums(t, got, expected)
	})

	t.Run("sum of tails of empty and full slice", func(t *testing.T) {
		got 		:= SumAllTail([]int{}, []int{2,3,3})
		expected	:= []int{0,6}

		checkSums(t, got, expected)
	})
}