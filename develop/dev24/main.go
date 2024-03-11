package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными
// параметрами x,y и конструктором.

type Point struct {
	x float64
	y float64
}

func  distance(a, b *Point) float64 {
	return math.Sqrt(((b.x - a.x) * (b.x - a.x)) + ((b.y - a.y) * (b.y - a.y)))
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func main() {
	a := NewPoint(5, 3)
	b := NewPoint(10, 7)

	result := distance(a, b)

	fmt.Println(result)
}
