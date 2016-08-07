package main

import (
	"fmt"
	"sort"
)

type people []string

//Sort package is being used here cause Interface interface is being implemented
func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	studyGroup := people{"Zeno", "John", "Ai", "Jenny"}
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	//exercise 2: sort this string s:= {"Zeno", "John", "Ai", "Jenny"}
	s := []string{"Zeno", "John", "Ai", "Jenny"}
	sort.StringSlice(s).Sort()
	fmt.Println(s)
	fmt.Println("Is string s sorted? ", sort.StringsAreSorted(s))
}
