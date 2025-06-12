package main

import "fmt"

type Vertex struct {
	x int
	y int
}

var (
	v1 = Vertex{1, 2}  // has type vetex
	v2 = Vertex{x: 1}  //y:0 is implicit
	v3 = Vertex{}      // x:0 and y:0
	g  = &Vertex{1, 3} //has type *Vetex
)

func main() {
	v := Vertex{1, 2}
	p := &v
	p.x = 1e9

	fmt.Println(v)
	fmt.Println(v1, v2, v3, g)
}
