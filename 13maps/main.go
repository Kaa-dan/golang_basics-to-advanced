package main

import "fmt"

type Vertex struct {
	lat, log float64
}

var m map[string]Vertex

func main() {

	//map should be comparable

	m = make(map[string]Vertex)

	m["Bell labs"] = Vertex{4.22, -73442.83}

	fmt.Println(m["Bell labs"])

	//map literal

	var b = map[string]Vertex{
		"Bell": {5.4345, 254.455},
		"ring": {324, 24},
	}
	fmt.Println(b)

	//never use var for initialising a map // check how
}
