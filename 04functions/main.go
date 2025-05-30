package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// function can return multiple value
func swap(x, y string) (string, string) {
	return x, y
}

//implcit return

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main() {
	fmt.Println(add(1, 2))

	a, b := swap("hello", "world")

	fmt.Println(a, b)

	fmt.Println(split(25))
}
