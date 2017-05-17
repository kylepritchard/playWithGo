package main

import "fmt"

func main() {
	slice := make([]string, 2, 10) //Makes a slice of referenced type so is available outside main function
	fmt.Printf("%T \n", slice)
	addToSlice(slice)
	slice[1] = "Pritchard"
	fmt.Println(slice)
}

func addToSlice(s []string) {
	s[0] = "Kyle" //sets index value '0' of the slice created in main function
	fmt.Println(s[0])
}
