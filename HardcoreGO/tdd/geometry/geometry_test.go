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

	t.Run("testing area of rectangle", func(t *testing.T) {
		rect 		:= Rectangle{10.0, 10.0}
		got 		:= rect.Area()
		expected	:= 100.0

		if got != expected{
			t.Errorf("Got %.2f but expected %.2f", got, expected)
		}
	})

	t.Run("testing area of circle", func(t *testing.T) {
		cir			:= Circle{10.0}
		got 		:= cir.Area()
		expected	:= 314.1592653589793

		if got != expected{
			t.Errorf("Got %0.2f but expected %0.2f", got, expected)
		}
	})


}