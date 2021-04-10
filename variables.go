package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var name string
	// var a uint8
	var a uint32
	a = 4444
	a = a + 10000000
	fmt.Println(a)
	name = "Milind"
	fmt.Println(name)

	// getrting the type of the variable
	num2 := 6.09
	fmt.Printf("%T", num2)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("types something ")
	scanner.Scan()
	input, error := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println(error)
	fmt.Println("You Typed ", 2021-input)

}
