package Shaper
import "testing"

func TestRectPerimeter(t *testing.T)	{
	var s Shaper
	r := Rect{5,3}
	s = r
	var rectPerimeter, wantRectPerimeter float64
	rectPerimeter = s.Perimeter()
	wantRectPerimeter = 16
	if rectPerimeter != wantRectPerimeter	{
		t.Errorf("Expected %.2f, got %.2f",wantRectPerimeter,rectPerimeter)
	}	
}

func TestCirclePerimeter(t *testing.T)	{
	var s Shaper
	c := Circle{5}
	s = c
	var circlePerimeter, wantCirclePerimeter float64
	circlePerimeter = s.Perimeter()
	wantCirclePerimeter = 31.42
	if circlePerimeter > wantCirclePerimeter	{
		t.Errorf("Expected %.2f, got %.2f",wantCirclePerimeter,circlePerimeter)
	}	
}