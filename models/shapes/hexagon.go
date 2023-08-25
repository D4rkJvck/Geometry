package models

import "math"

type Hexagon struct {
	Side float64
}

/////////////////////////////////////////////////////////
//-> Methods	|
////////////	V

func (h *Hexagon) Area() float64 {
	return math.Pow(h.Side, 2) * (3 * math.Sqrt(3)) / 2
}

func (h *Hexagon) Perimeter() float64 {
	return h.Side * 6
}