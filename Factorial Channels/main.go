package main

import "fmt"

func main() {

	f := make(chan int)
	var factorial = 1

	go func() {
		for i := 4; i > 0; i-- {
			f <- i
		}
		close(f)
		fmt.Println(factorial)
	}()

	for n := range f {
		fmt.Println(n)
		factorial *= n
	}
}
