package main

import "fmt"

func main() {
	i := 7
	// Will display address of i
	// Basics
	// & = address of
	// * = value at
	fmt.Println(&i)
	inc(&i)
	fmt.Println(i)
}

func inc(x *int) {

	*x++
}

// output
// 0xc0000180e8
// 8
