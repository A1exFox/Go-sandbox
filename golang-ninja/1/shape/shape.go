package shape

import "math"

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}
type ShapeWithArea interface {
	Area() float32
}
type ShapeWithPerimeter interface {
	Perimeter() float32
}

type Square struct {
	sideLength float32
}

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}
func (s Square) Perimeter() float32 {
	return s.sideLength * 4
}
func NewSquare(length float32) Square {
	return Square{
		sideLength: length,
	}
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}
func (c Circle) Perimeter() float32 {
	return c.radius * math.Pi * 2
}
func NewCircle(radius float32) Circle {
	return Circle{radius: radius}
}
