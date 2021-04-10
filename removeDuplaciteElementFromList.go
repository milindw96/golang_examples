// removing dupliacte element from the list
package main

import "fmt"

func main() {

	// range
	var a []int = []int{1, 2, 4, 5, 6, 4, 2, 7}

	var newSlice []int = []int{}
	test := len(a)

	for index, element := range a {
		flag := 0
		for index1, element1 := range a {

			if element == element1 && index1 < index {
				flag = 1
			}
		}
		if flag == 0 {
			fmt.Println("flag = ", flag)

			// fmt.Println(element)
			newSlice = append(newSlice, element)
		}

		// fmt.Println("flag= ", flag, index, " ", element)
	}
	fmt.Println(newSlice)
	test = test - 1
}
