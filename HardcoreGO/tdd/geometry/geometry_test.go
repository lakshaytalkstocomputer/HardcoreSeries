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
		name string
		shape Shape
		expected float64
	}{
		{"Rectangle",Rectangle{12,6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.expected{
				t.Errorf("got %g want %g in %#v", got, tt.expected, tt.shape)
			}
		})
	}
}