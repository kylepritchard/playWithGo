package main

import (
	"fmt"
)

func main() {

	var a int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&a)

	divideAndEven := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}

	fmt.Println(divideAndEven(a))
}
