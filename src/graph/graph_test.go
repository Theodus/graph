package main

import (
    "testing"
)

func TestGraph(t *testing.T) {
    g, err := NewGraph("Test Graph", "X", "Y")
    if err != nil {
        t.Fatal(err)
    }
    g.AddPoint(1.0, 5.7)
    g.AddPoint(2.0, 7.8)
    g.AddPoint(3.0, 10.1)
    g.Save(400, 300, "src/graph/test/testgraph1.pdf")
}
