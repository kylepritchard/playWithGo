package main

import "fmt"

func main() {

	/*
		func main() {
			foo(1, 2)
			foo(1, 2, 3)
			aSlice := []int{1, 2, 3, 4}
			foo(aSlice...)
			foo()
		}
	*/

	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()

}

func foo(numbers ...int) {
	fmt.Println("The numbers sent to foo() are:", numbers)
}
