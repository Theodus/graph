package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Data struct {
	Name   string `json:"name"`
	X      string `json:"x"`
	Y      string `json:"y"`
	W      int    `json:"w"`
	H      int    `json:"h"`
	File   string `json:"file"`
	Curves []struct {
		Name string    `json:"name"`
		RGB  []uint8   `json:"rgb"`
		X    []float64 `json:"x"`
		Y    []float64 `json:"y"`
	} `json:"curves"`
}

func Parse(filename string) error {
	var d Data
	fmt.Println("Reading", filename)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, &d); err != nil {
		return err
	}
	fmt.Println("Graphing...")
	g, err := NewGraph(d.Name, d.X, d.Y)
	for _, e := range d.Curves {
		c := &Curve{Name: e.Name, RGB: e.RGB}
		for i := range e.X {
			c.AddPoint(e.X[i], e.Y[i])
		}
		g.AddCurve(*c)
	}
	fn := d.File
	if err = g.Save(fn, d.W, d.H); err != nil {
		return err
	}
	fmt.Println("Saving file to", fn)
	fmt.Println("Done!")
	return nil
}
