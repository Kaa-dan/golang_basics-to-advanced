package main

import "fmt"

// generics
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {

	si := []int{10, 20, 30}
	fmt.Println(Index(si, 20))

	ss := []string{"nithin", "raj"}
	fmt.Println(Index(ss, "raj"))
}
