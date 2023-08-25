package models

import "math"

type Square struct {
	Side float64
}

////////////////////////////////////////
//-> Methods	|
////////////	V

func (s *Square) Area() float64 {
	return math.Pow(s.Side, 2)
}

func (s *Square) Perimeter() float64 {
	return s.Side * 4
}

