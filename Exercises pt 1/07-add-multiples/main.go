package main

import "fmt"

func main() {

	sum := 0

	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		} else {
			continue
		}
	}

	fmt.Println("The sum of all mutiples of 3 & 5 up to 1000 =", sum)
}
