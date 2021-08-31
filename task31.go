package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func DistanceBetweenPoints(a, b Point) float64 {
	return math.Sqrt(math.Pow(float64(b.x-a.x), 2) + math.Pow(float64(b.y-a.y), 2))
}

func main() {
	a := NewPoint(3, -4)
	b := NewPoint(1, -6)
	fmt.Println(DistanceBetweenPoints(a, b))
}
