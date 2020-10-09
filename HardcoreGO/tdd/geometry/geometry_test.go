package geometry

import "testing"

func TestPerimeter(t *testing.T){
	got 		:= Perimeter(10.0, 10.0)
	expected 	:= 40.0

	if got != expected{
		t.Errorf("Got %0.2f Expected %.2f", got, expected)
	}

}