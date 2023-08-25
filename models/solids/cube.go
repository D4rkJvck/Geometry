package models

import "math"

type Cube struct {
	Side float64
}

///////////////////////////////////////
//-> Methods	|
////////////	V

func (c *Cube) LateralArea() float64 {
	return math.Pow(c.Side, 2) * 4
}

func (c *Cube) Area() float64 {
	return math.Pow(c.Side, 2) * 6
}

func (c *Cube) Volume() float64 {
	return math.Pow(c.Side, 3)
}
