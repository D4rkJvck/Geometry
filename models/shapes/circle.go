package models

import "math"

type Circle struct {
	Radius float64
}

///////////////////////////////////////////
//-> Methods	|
////////////	V

func (c *Circle) Diameter() float64 {
	return c.Radius * 2
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c *Circle) Perimeter() float64 {
	return c.Diameter() * math.Pi
}
