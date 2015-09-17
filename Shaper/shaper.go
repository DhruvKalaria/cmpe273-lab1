package Shaper
import	(
	"fmt"
	"math"
)

type Shaper interface	{
	
	Area() float64
	Perimeter() float64
}

type Rect struct	{
	length, breadth float64
}

func (r Rect) Area() float64	{
	return r.length*r.breadth
}

func (r Rect) Perimeter() float64	{
	return 2*(r.length+r.breadth)
}

type Circle struct	{
	radius float64
}

func (c Circle) Area() float64	{
	return math.Pi*c.radius*c.radius
}

func (c Circle) Perimeter() float64	{
	return 2*math.Pi*c.radius
}

func main()	{
	var s Shaper
	r := Rect{5,3}
	c := Circle{5}
	s = r
	fmt.Printf("Area of Rectangle with Length=%.2f and Breadth=%.2f is:%.2f\n",r.length,r.breadth,s.Area())	
	fmt.Printf("Perimeter of Rectangle with Length=%.2f and Breadth=%.2f is:%.2f\n",r.length,r.breadth,s.Perimeter())	
	
	s = c
	fmt.Printf("Area of Circle with Radius=%.2f is:%.2f\n",c.radius,s.Area())	
	fmt.Printf("Perimeter of Circle with Radius=%.2f is:%.2f\n",c.radius,s.Perimeter())
}