package main

import "fmt"

func main() {
	fmt.Println(findBiggest(14, 26, 4, 52, 31, 2, 13456))
}

// numbers is variadic parameter
func findBiggest(numbers ...int) int {
	var biggest int
	//iterate over numbers
	for _, v := range numbers {
		if v > biggest {
			biggest = v
		}
	}

	return biggest
}
