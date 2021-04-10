package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apple":     5,
		"pineapple": 3,
		"2":         2,
		"5":         5,
	}
	fmt.Print(mp)
	mp["apple"] = 7
	// adding new value with in map
	mp["apple2"] = 9090
	fmt.Println(mp)

	// create new empty map
	make_newmap := make(map[string]int)
	fmt.Println(make_newmap)
}
