package main

import "fmt"

type Point struct {
	x int32
	y int32
}
type Circle struct {
	radius float64
	center *Point
}

func main() {
	// Struct  are used create own coustome behaviour/type
	var p1 Point = Point{1, 2}
	var p2 Point = Point{20, 9}
	fmt.Println(p1.x, p2.x)
	fmt.Println(p1.y, p2.y)
	p1.x = 71
	fmt.Println(p1.x, p2.x)
	p3 := Point{x: 3}
	fmt.Println(p3.x, p3.y)
	c1 := Circle{3.46, &Point{3, 5}}
	fmt.Println(c1, c1.center.x)
}
