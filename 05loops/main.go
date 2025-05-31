package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}

	//while loop

	for sum > 10 {
		fmt.Println("nithin")
	}

	//infinite loop
	for {
		if sum > 100 {
			break
		}
		sum += 1
	}
	fmt.Println(sum)
}
