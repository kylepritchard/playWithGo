// Use https://godoc.org/sort to sort the following:

// (1)
// type people []string
// studyGroup := people{"Zeno", "John", "Al", "Jenny"}

// (2)
// s := []string{"Zeno", "John", "Al", "Jenny"}

// (3)
// n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

// Also sort the above in reverse order

package main

import (
	"fmt"
	"sort"
)

type people []string //custom type

// need to reference to type people to use Sort
func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {

	peeps := people{"Kyle", "Robyn", "Kady", "Ross", "Max", "Harley"}
	sort.Sort(peeps)
	fmt.Println("1: Sort alphabetically", peeps)

	s := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(sort.StringSlice(s))
	fmt.Println("2: Sort slice of Strings", s)

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Sort(sort.IntSlice(n))
	fmt.Println("3: Integers sorted", n)

	//Reverse
	sort.Sort(sort.Reverse(peeps))
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	sort.Sort(sort.Reverse(sort.IntSlice(n)))

	fmt.Println("4: People in reverse", peeps)
	fmt.Println("5: Strings in reverse", s)
	fmt.Println("6: Numbers in reverse", n)
}
