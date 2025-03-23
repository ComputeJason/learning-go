package structs

import "math"


type Shape interface{
	getPerimeter()float64;
	getArea()float64;
}

type Rectangle struct {
	height float64;
	width float64;
}

func (r Rectangle) getPerimeter() float64 {
	return r.height * 2 + r.width * 2 
}

func (r Rectangle) getArea() float64 {
	return r.height * r.width
}

type Circle struct {
	radius float64
}

func (c Circle) getPerimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}

type Triangle struct {
	base float64
	height float64
}

func (t Triangle) getPerimeter() float64 {
	return 0
}

func (t Triangle) getArea() float64 {
	return (1.0/2) * t.base * t.height
}



