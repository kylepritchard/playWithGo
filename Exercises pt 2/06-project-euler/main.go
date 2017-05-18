// A palindromic number reads the same both ways. The largest palindrome
// made from the product of two 2-digit numbers is 9009 = 91 × 99.

// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {

	//work out all the multiples from 100 to 999 (x * y = z)

	var z int
	var z_string string
	var biggest [3]int //save biggest palindrome info in an array

	timeStart := time.Now()

	for x := 100; x <= 999; x++ {
		for y := 100; y <= 999; y++ {
			z = x * y
			z_string = strconv.Itoa(z) //convert to a string
			if checkPalindrome(z_string) && z > biggest[2] {
				biggest[0] = x
				biggest[1] = y
				biggest[2] = z
			}
		}
	}

	fmt.Println("The biggest palindrome is", biggest[2], "which is", biggest[0], "=", biggest[1])

	timeElapsed := time.Since(timeStart)
	log.Printf("Finding palindromes took %s", timeElapsed)

}

func checkPalindrome(s string) bool {
	var length int = len(s)
	var fromEnd int = 0

	for i := 0; i < (length / 2); i++ {
		fromEnd = (length - 1) - i //length of string is 1 higher than index
		// if ith character of string = ith character from end ie 12345 if 1 = 5 or 2 = 4....
		if s[i] != s[fromEnd] {
			return false // return false and stop checking
		}
	}
	return true // its a palindrome!
}
