package main

import "fmt"

func main() {
	m := 3

	switch {
	case m > 3, m < 3:
		fmt.Println("done")
	default:
		fmt.Println("heY")
	}

	fmt.Println(m)
}
