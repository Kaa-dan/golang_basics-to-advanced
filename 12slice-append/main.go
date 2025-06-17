package main

import "fmt"

func main() {

	var s []int

	//append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	//the slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	//we can add more than one element at a time

	s = append(s, 2, 3, 4, 5)
	printSlice(s)
	//here range will return two values index and copy of the value
	for i, v := range s {
		fmt.Println(i, v)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
