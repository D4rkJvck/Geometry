package models

import "math"

type Triangle struct {
	Side_A float64
	Side_B float64
	Side_C float64
}

////////////////////////////////////////////////////////////////////////////////////////
//-> Methods							|
////////////							V

func (t *Triangle) SemiPerimeter() float64 {
	return (t.Side_A + t.Side_B + t.Side_C) / 2
}

func (t *Triangle) Area() float64 {
	semi := t.SemiPerimeter()
	return math.Sqrt(semi * (semi - t.Side_A) * (semi - t.Side_B) * (semi - t.Side_C))
}

func (t *Triangle) Perimeter() float64 {
	return t.SemiPerimeter() * 2
}
