package main

import (
	"fmt"
	"strings"
)

func main() {

	names := [4]string{"nithin", "raja", "vadhu", "vatt"}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	//slice literal
	q := []int{1, 2, 3, 4, 5}
	fmt.Println(q)

	s := []struct {
		i int
		b bool
	}{{2, true}, {3, false}}

	fmt.Println(s)

	k := []int{1, 2, 3, 4}

	k = k[1:3]

	k = k[:2]

	k = k[1:]

	fmt.Println(k)
	printSlice(k)

	// 0 value

	var z []int

	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
	//using make
	mk := make([]int, 5)
	printSlice(mk)

	// two dimensional slices

	board := [][]string{{"nithin", "raj"}, {"nithin", "raj"}}

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], ""))
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
