package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func Init(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) DistanceTo(point *Point) float64 {
	//Теорема пифагора :)
	return math.Sqrt(math.Pow(math.Abs(p.x-point.x), 2) + math.Pow(math.Abs(p.y-point.y), 2))
}

func main() {
	point1 := Init(0, 0)
	point2 := Init(10, 10)
	fmt.Println(point1.DistanceTo(point2))

	point1 = Init(-10, -10)
	fmt.Println(point1.DistanceTo(point2))
}
