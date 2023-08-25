package models

import "math"

type Pentagon struct {
	Side float64
}

//////////////////////////////////////////////////////
//-> Methods	|
////////////	V

func (p *Pentagon) Area() float64 {
	tan := math.Tan((54 * math.Pi) / 180)
	return math.Pow(p.Side, 2) * ((p.Side * tan) / 4)
}

func (p *Pentagon) Perimeter() float64 {
	return p.Side * 5
}
