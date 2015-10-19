package main

type Curve struct {
	Name   string
	RGB    []uint8
	Points Points
}

func (c *Curve) AddPoint(x, y float64) {
	c.Points = append(c.Points, Point{x, y})
}
