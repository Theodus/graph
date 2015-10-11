package main

import(
    "github.com/gonum/plot/plotutil"
    "github.com/gonum/plot"
    "github.com/gonum/plot/vg"
)

type Graph struct {
    Plot     *plot.Plot
    Points   Points
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

func (g *Graph) AddPoint(x, y float64) {
    g.Points = append(g.Points, Point{x, y})
}

func (g *Graph) Save(w, h int, filename string) error {
    err := plotutil.AddLinePoints(g.Plot, g.Points)
    if err != nil {
        return err
    }
    return g.Plot.Save(vg.Length(w), vg.Length(w), filename)
}

//todo add graph key and point shapes