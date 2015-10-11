package main

import (
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"image/color"
)

type Graph struct {
	Plot   *plot.Plot
	Curves []Curve
}

func NewGraph(title, x, y string) (g *Graph, err error) {
	p, err := plot.New()
	if err != nil {
		return nil, err
	}
	p.Title.Text = title
	p.X.Label.Text = x
	p.Y.Label.Text = y
	g = &Graph{Plot: p}
	return g, err
}

func (g *Graph) AddCurve(c Curve) {
	g.Curves = append(g.Curves, c)
}

func (g *Graph) Save(filename string, w, h int) error {
	g.Plot.Add(plotter.NewGrid())
	for i, e := range g.Curves {
		lpLine, lpPoints, err := plotter.NewLinePoints(e.Points)
		if err != nil {
			return err
		}
		switch i {
		case 0:
			color := color.RGBA{R: 50, G: 50, B: 220, A: 255}
			lpLine.LineStyle.Color = color
			lpPoints.Color = color
		case 1:
			color := color.RGBA{R: 220, G: 50, B: 50, A: 255}
			lpLine.LineStyle.Color = color
			lpPoints.Color = color
		case 2:
			color := color.RGBA{R: 50, G: 220, B: 50, A: 255}
			lpLine.LineStyle.Color = color
			lpPoints.Color = color
		case 3:
			color := color.RGBA{R: 220, G: 50, B: 220, A: 255}
			lpLine.LineStyle.Color = color
			lpPoints.Color = color
		case 4:
			color := color.RGBA{R: 50, G: 220, B: 220, A: 255}
			lpLine.LineStyle.Color = color
			lpPoints.Color = color
		case 5:
			color := color.RGBA{R: 220, G: 220, B: 50, A: 255}
			lpLine.LineStyle.Color = color
			lpPoints.Color = color
		}
		g.Plot.Add(lpLine, lpPoints)
		g.Plot.Legend.Add(e.Name, lpLine, lpPoints)
	}
	return g.Plot.Save(vg.Length(w), vg.Length(h), filename)
}
