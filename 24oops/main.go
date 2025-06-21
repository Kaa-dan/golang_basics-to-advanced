package main

import "fmt"

type Number int

func (n *Number) is_even() bool {
	fmt.Println(&n)
	fmt.Println(n)
	fmt.Println(*n)
	*n = 20
	return *n%2 == 0
}
func main() {
	var v Number = 10

	fmt.Println(v.is_even())
	fmt.Println(v.is_even())
	fmt.Println(v)
}
