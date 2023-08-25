package models

type Rectangle struct {
	Lenght float64
	Width  float64
}

///////////////////////////////////////////
//-> Methods	|
////////////	V

func (r *Rectangle) Area() float64 {
	return r.Lenght * r.Width
}

func (r *Rectangle) Perimeter() float64 {
	return (r.Lenght * 2) + (r.Width * 2)
}
