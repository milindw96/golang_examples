package main

import "fmt"

// mutable datatype = changable
// unmutable Datatype = not changable
func main() {
	// unmutable Datatype Example
	// same behaviour is expected in array

	var x int = 5
	y := x
	x = 7
	fmt.Println(x, y)
	// mutable datatype Example
	var a []int = []int{4, 5, 6}
	b := a
	b[0] = 9090
	fmt.Println(a, b)
	/*
		Output will be
		[9090 5 6] [9090 5 6]
		Reason a and b is pointing to same slice
		same behaviour is expected in map also
	*/
}
