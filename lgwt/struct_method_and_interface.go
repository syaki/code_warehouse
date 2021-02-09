package main

import "math"

type Shape interface {
	area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (c Triangle) area() float64 {
	return (c.Base * c.Height) * 0.5
}

func perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
