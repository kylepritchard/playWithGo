package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {

	book := "http://www.gutenberg.org/files/1342/1342-0.txt"
	// book := "123"

	res, err := http.Get(book)
	// check if there has been an error getting from interweb
	if err != nil {
		log.Fatal(err) //end the program with exit status 1 and give error
	}

	// Make buckets map

	buckets := make(map[int]map[string]int) //map of [1:map[string]int 2:map....    ]

	// make 12 buckets
	for i := 0; i < 12; i++ {
		buckets[i] = make(map[string]int) // make map [string]int in each 'bucket'
	}

	// Scan the response from http.Get(book)
	scanner := bufio.NewScanner(res.Body) //scan to a buffer named scanner
	defer res.Body.Close()                //make sure the response is closed at end of function
	scanner.Split(bufio.ScanWords)        //split into words

	// Loop over the words
	for scanner.Scan() {
		word := scanner.Text() // get each word and save as word
		n := hashBucket(word)  //hash word to get bucket number, n is the bucket number returned by function
		buckets[n][word]++     //counts time the word appears - starts at 0 and adds by one each time word appears
	}

	// Print words in a bucket
	// for k, v := range buckets[6] {
	// 	fmt.Println(v, " \t- ", k)
	// }

	// search for words in the book...
	for i := 0; i < 12; i++ {
		if val, ok := buckets[i]["hello"]; ok {
			fmt.Println(val)
		}

	}

}

// find out the hash bucket
// return the bucket number
func hashBucket(word string) int {
	var sum int // add the value of each character in word
	for _, v := range word {
		sum += int(v)
	}
	return sum % 12

}
