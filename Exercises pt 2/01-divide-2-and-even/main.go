package main

import "fmt"

func main() {

	var a int
	fmt.Print("Enter an integer to process: ")
	fmt.Scan(&a)
	half, even := divideByTwo(a)
	fmt.Println("Half of the integer =", half)
	fmt.Println("Is the number even? - ", even)
}

func divideByTwo(n int) (int, bool) {

	half := n / 2
	even := false
	if n%2 == 0 {
		even = true
	}

	return half, even
}
