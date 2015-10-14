package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Data struct {
  Name   string
	X      string
	Y      string
	W      int
	H      int
	File   string
	Curves []struct {
		Name string
		X    []float64
		Y    []float64
	}
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
		c := &Curve{Name: e.Name}
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
