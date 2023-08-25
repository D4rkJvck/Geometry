package models

type Prism struct {
	Side_A float64
	Side_B float64
	Side_C float64
}

/////////////////////////////////////////////////////////////////
//-> Methods	|
////////////	V

func (p *Prism) LateralArea() float64 {
	return (2 * p.Side_A * p.Side_B) + (2 * p.Side_A * p.Side_C)
}

func (p *Prism) Area() float64 {
	return p.LateralArea() + (2 * p.Side_B * p.Side_C)
}

func (p *Prism) Volume() float64 {
	return p.Side_A * p.Side_B * p.Side_C
}
