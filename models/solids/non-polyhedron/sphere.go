package models

import "math"

type Sphere struct {
	Radius float64
}

///////////////////////////////////////////
//-> Methods	|
////////////	V

func (s *Sphere) Area() float64 {
	return 4 * math.Pi * math.Pow(s.Radius, 2)
}

func (s *Sphere) Volume() float64 {
	return (4 * math.Pi * math.Pow(s.Radius, 3)) / 3
}
