package main

import (
	"fmt"
	"sort"
)

type intList []int

func (n intList) Len() int           { return len(n) }
func (n intList) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n intList) Less(i, j int) bool { return n[i] < n[j] }

func main() {
	n := []int{7, 5, 8, 9, 19, 12, 32, 3}
	sort.Ints(n)
	fmt.Println("Sortign in increasing orders ", n)

	// Reverse order of the data
	m := intList{7, 5, 8, 9, 19, 12, 32, 3}
	fmt.Print("type of m before doing the conversion to intSlice is: ")
	fmt.Printf("%T \n", m)
	//sort.IntSlice(m)
	fmt.Print("type of sort.IntSlice(m) is: ")
	fmt.Printf("%T \n", sort.IntSlice(m))
	fmt.Print("type of m after doing the conversion to intSlice is: ")
	fmt.Printf("%T \n", m)

	sort.Sort(sort.Reverse(sort.IntSlice(m)))
	fmt.Println("Sorting in reverse orders ", m)
	fmt.Print("type of m after doing the reverse is: ")
	fmt.Printf("%T \n", m)

	fmt.Print("type of (sort.Reverse(sort.IntSlice(m))) is: ")
	fmt.Printf("%T \n", (sort.Reverse(sort.IntSlice(m))))
}
