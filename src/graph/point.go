package main

type Point struct {
    X float64
    Y float64
}

type Points []Point

func (p Points) Len() int {
    return len([]Point(p))
}

func (p Points) XY(i int) (float64, float64) {
    return p[i].X, p[i].Y
}
