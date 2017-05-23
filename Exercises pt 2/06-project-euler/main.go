// A palindromic number reads the same both ways. The largest palindrome
// made from the product of two 2-digit numbers is 9009 = 91 × 99.

// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	//work out all the multiples from 100 to 999 (x * y = z)
	iterations := 0
	var z int
	var zString string
	var biggest [3]int //save biggest palindrome info in an array

	timeStart := time.Now()

	for x := 100; x <= 999; x++ {
		for y := 100; y <= 999; y++ {
			iterations++
			z = x * y
			zString = strconv.Itoa(z) //convert to a string
			if checkPalindrome(zString) && z >= biggest[2] {
				biggest[0] = x
				biggest[1] = y
				biggest[2] = z
			}
		}
	}
	fmt.Println("Number of iterations done:", iterations)
	fmt.Println("The biggest palindrome is", biggest[2], "which is", biggest[0], "*", biggest[1])

	timeElapsed := time.Since(timeStart)
	fmt.Printf("Finding palindromes took %s \n", timeElapsed)

}

func checkPalindrome(s string) bool {
	var length = len(s)
	var fromEnd int

	for i := 0; i < (length / 2); i++ {
		fromEnd = (length - 1) - i //length of string is 1 higher than index
		// if ith character of string = ith character from end ie 12345 if 1 = 5 or 2 = 4....
		if s[i] != s[fromEnd] {
			return false // return false and stop checking
		}
	}
	return true // its a palindrome!
}
