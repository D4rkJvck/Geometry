package models

import "math"

type Trapezoid struct {
	Base_A float64
	Base_B float64
	Height float64
}

////////////////////////////////////////////////////////////////////
//-> Methods	|
////////////	V

func (t *Trapezoid) SemiPerimeter() float64 {
	diff := t.Base_B - t.Base_A
	return math.Sqrt(math.Pow((diff/2), 2) + math.Pow(t.Height, 2))
}

func (t *Trapezoid) Area() float64 {
	return ((t.Base_A + t.Base_B) * t.Height) / 2
}

func (t *Trapezoid) Perimeter() float64 {
	return (t.SemiPerimeter() * 2) + t.Base_A + t.Base_B
}
