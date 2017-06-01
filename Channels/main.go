package main

import (
	"fmt"
	"time"
)

func main() {

	out := make(chan int)
	sum := 0

	timeStart := time.Now()
	go func() {
		for i := 0; i < 1000000; i++ {
			out <- i
		}
		close(out)
		// println("goroutine finished")

	}()

	for n := range out {
		sum += n
	}

	fmt.Println(sum)
	time := time.Since(timeStart)
	fmt.Println(time)

}

// func main() {
// 	// Set up the pipeline and consume the output.
// 	for n := range sq(sq(gen(2, 3))) {
// 		fmt.Println(n) // 16 then 81
// 	}
// }

// func gen(nums ...int) chan int {
// 	out := make(chan int)
// 	go func() {
// 		for _, n := range nums {
// 			out <- n
// 		}
// 		close(out)
// 	}()
// 	fmt.Printf("%T - %s\n", out, out)
// 	return out
// }

// func sq(in chan int) chan int {
// 	out := make(chan int)
// 	go func() {
// 		for n := range in {
// 			out <- n * n
// 		}
// 		close(out)
// 	}()
// 	return out
// }
