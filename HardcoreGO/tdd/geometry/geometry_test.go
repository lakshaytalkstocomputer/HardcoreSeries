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

	areaTests := []struct{
		shape Shape
		expected float64
	}{
		{Rectangle{12,6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests{
		got := tt.shape.Area()

		if got != tt.expected{
			t.Errorf("got %g want %g in %#v", got, tt.expected, tt.shape)
		}
	}

}