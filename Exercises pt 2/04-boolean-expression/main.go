package main

import "fmt"

func main() {

	// What's the value of this expression: (true && false) || (false && true) || !(false && false)?
	expression := (true && false) || (false && true) || !(false && false)
	fmt.Println(expression)

}
