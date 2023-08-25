package models

import "math"

type Pyramid struct {
	Side   float64
	Height float64
}

///////////////////////////////////////////////////////////////////////////////
//-> Methods	|
////////////	V

func (p *Pyramid) LateralArea() float64 {
	return 2 * p.Side * math.Sqrt((math.Pow(p.Height, 2) + math.Pow(p.Side/2, 2)))
}

func (p *Pyramid) Area() float64 {
	return p.LateralArea() + math.Pow(p.Side, 2)
}

func (p *Pyramid) Volume() float64 {
	return math.Pow(p.Side, 2) * (p.Height / 3)
}
