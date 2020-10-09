package geometry

import "testing"

func TestPerimeter(t *testing.T){
	got 		:= Perimeter(10.0, 10.0)
	expected 	:= 40.0

	if got != expected{
		t.Errorf("Got %.2f Expected %.2f", got, expected)
	}

}


func TestArea(t *testing.T){
	got			:= Area(10.0, 10.0)
	expected	:= 100.0

	if got != expected{
		t.Errorf("Got %.2f but expected %.2f", got, expected)
	}
}