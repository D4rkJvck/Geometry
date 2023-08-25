package models

import "math"

type Cone struct {
	Radius float64
	Height float64
}

///////////////////////////////////////////////////////////////////////////////////////
//-> Methods	|
////////////	V

func (c *Cone) LateralArea() float64 {
	return c.Radius * math.Pi * math.Sqrt((math.Pow(c.Radius, 2) + math.Pow(c.Height, 2)))
}

func (c *Cone) Area() float64 {
	return c.LateralArea() + (math.Pi * math.Pow(c.Radius, 2))
}

func (c *Cone) Volume() float64 {
	return (math.Pi * math.Pow(c.Radius, 2) * c.Height) / 3
}
