package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //send sum to c
}

func main() {
	s := []int{1, 2, 3, 4, 5}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c //receive from c

	fmt.Println(x, y, x+y)

	//buffer

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	chanel := make(chan int, 10)
	go fibonacci(cap(chanel), c)

	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
