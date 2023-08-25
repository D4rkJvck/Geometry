package models

type Cuboid struct {
	Side_A float64
	Side_B float64
	Side_C float64
}

/////////////////////////////////////////////////////////////////
//-> Methods	|
////////////	V

func (c *Cuboid) LateralArea() float64 {
	return (2 * c.Side_A * c.Side_B) + (2 * c.Side_A * c.Side_C)
}

func (c *Cuboid) Area() float64 {
	return c.LateralArea() + (2 * c.Side_B * c.Side_C)
}

func (c *Cuboid) Volume() float64 {
	return c.Side_A * c.Side_B * c.Side_C
}
