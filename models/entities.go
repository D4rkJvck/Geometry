package models

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Solid interface {
	Area() float64
	Volume() float64
}
