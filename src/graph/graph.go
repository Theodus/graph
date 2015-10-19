package main

import (
	"errors"
	"image/color"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
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
	for _, c := range g.Curves {
		lpLine, lpPoints, err := plotter.NewLinePoints(c.Points)
		if err != nil {
			return err
		}
		if len(c.RGB) != 3 {
			return errors.New("bad RGB")
		}
		color := color.RGBA{R: c.RGB[0], G: c.RGB[1], B: c.RGB[2], A: 255}
		lpLine.LineStyle.Color = color
		lpPoints.Color = color

		g.Plot.Add(lpLine, lpPoints)
		if c.Name != "" {
			g.Plot.Legend.Add(c.Name, lpLine, lpPoints)
		}
	}
	return g.Plot.Save(vg.Length(w), vg.Length(h), filename)
}
