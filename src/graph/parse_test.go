package main

import (
  "testing"
)

func TestParse(t *testing.T) {
  err := Parse("test/testgraph.json")
  if err != nil {
    t.Fatal(err)
  }
}