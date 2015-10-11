package main

import (
	"testing"
)

func TestGraph(t *testing.T) {
	g, err := NewGraph("Test Graph", "X", "Y")
	if err != nil {
		t.Fatal(err)
	}
	c := &Curve{Name: "Lonely Curve"}
	c.AddPoint(1.0, 5.7)
	c.AddPoint(2.0, 7.8)
	c.AddPoint(3.0, 10.1)
	g.Save("src/graph/test/testgraph1.pdf", 400, 300)
}
