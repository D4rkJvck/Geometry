package models

import "math"

type Cylinder struct {
	Radius float64
	Height float64
}

///////////////////////////////////////////////////////////////////
//-> Methods	|
////////////	V

func (c *Cylinder) LateralArea() float64 {
	return 2 * math.Pi * c.Radius * c.Height
}

func (c *Cylinder) Area() float64 {
	return c.LateralArea() + (2 * math.Pi * math.Pow(c.Radius, 2))
}

func (c *Cylinder) Volume() float64 {
	return math.Pi * math.Pow(c.Radius, 2) * c.Height
}
