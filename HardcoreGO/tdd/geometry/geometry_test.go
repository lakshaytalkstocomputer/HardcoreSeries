package geometry

import "testing"

func TestPerimeter(t *testing.T){
	t.Run("testing perimeter of rectangle", func(t *testing.T) {
		rect 		:= Rectangle{10.0, 10.0}
		got 		:= rect.Perimeter()
		expected 	:= 40.0

		if got != expected{
			t.Errorf("Got %.2f Expected %.2f", got, expected)
		}
	})

}


func TestArea(t *testing.T){

	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()

		got	:= shape.Area()

		if got != expected{
			t.Errorf("Got %.2f but expected %.2f", got, expected)
		}
	}


	t.Run("testing area of rectangle", func(t *testing.T) {
		rect 		:= Rectangle{10.0, 10.0}
		expected	:= 100.0

		checkArea(t, rect, expected)
	})

	t.Run("testing area of circle", func(t *testing.T) {
		cir			:= Circle{10.0}
		expected	:= 314.1592653589793

		checkArea(t, cir, expected)
	})


}