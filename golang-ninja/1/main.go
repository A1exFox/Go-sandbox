package main

import (
	"fmt"
	"golang-ninja/basic/shape"
)

func main() {
	square := shape.NewSquare(5)
	circle := shape.NewCircle(5)
	printShapeArea(square)
	printShapeArea(circle)
}

func printShapeArea(s shape.Shape) {
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter: ", s.Perimeter())
}
