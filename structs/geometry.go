// Notes

// Shape is an interface that defines the Area method
// In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.
// In this case, Rectangle and Circle both have Area() methods that return a Float64 value so they satisfy the Shape interface.

// passing a Rectangle to a function conveys intent more clearly then just values

// Method declaration syntax - func (receiverName ReceiverType) MethodName(args)
// method on the Rectangle struct

// Go convention is to use the first letter of the type as the receiver name
// in this case c for Circle

package structs

import "math"

// Shape is an interface that defines methods for geometric shapes
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle represents a rectangle with height and width
type Rectangle struct {
	Height, Width float64
}

// Area calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Perimeter calculates the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

// Circle represents a circle with a radius
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculates the perimeter of a circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Triangle represents a triangle with base and height
type Triangle struct {
	Base, Height float64
}

// Area calculates the area of a triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Perimeter calculates the perimeter of a triangle
// Assuming an equilateral triangle for simplicity
func (t Triangle) Perimeter() float64 {
	return 3 * t.Base
}
