package main

import (
	"testing"
	"fmt"
)
/*
func TestAverage(t *testing.T) {
  var v int
  v = Average(23)
  if v != 23 {
    t.Error("Expected 1.5, got ", v)
  }
}*/


func BenchmarkHello(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("hello")
    }
}
