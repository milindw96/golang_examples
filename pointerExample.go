package main

import "fmt"

// *string = expected pointer for the variable
func changeValue(str *string) {
	*str = "changed"
}

func changeValue2(str string) {
	str = "changed_2"
}

func main() {
	x := 7
	y := &x
	fmt.Println(&x)  //& =  address of
	fmt.Println(*&x) // * = vaule at that addredd
	*y = 90
	fmt.Println(x, y)

	toChange := "hello"
	fmt.Println(toChange)
	changeValue(&toChange)
	fmt.Println(toChange)

	toChange2 := "Hello"
	fmt.Println(toChange2)
	changeValue2(toChange2)
	fmt.Println(toChange2)

}
