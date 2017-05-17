package main

import "fmt"

func main() {

	var bigNumber int
	var smallNumber int

	fmt.Print("Enter the big number: ")
	fmt.Scan(&bigNumber)
	fmt.Print("Enter the small number: ")
	fmt.Scan(&smallNumber)

	remainder := bigNumber % smallNumber

	fmt.Println("The remainder is:", remainder)

}
