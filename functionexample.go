package main

import "fmt"

func funcTestng(a, b int) (int, string) {
	defer fmt.Println("i am defering")
	// defer The deferred call's arguments are evaluated immediately,
	// but the function call is not executed until the surrounding function returns.
	// defer will execute at the end

	fmt.Println("testing")
	return a + b, "hello"
}

func main() {
	output, output2 := funcTestng(2, 4)
	fmt.Println(output, output2)

}
