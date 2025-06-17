package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Vertex struct {
	X, Y float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  //a MyFloat implements Abser
	a = &v //a *Vertex implements Abser

	//in the following line,  i a Vertex (not *Vertex)
	// and does Not implement Absert.

	a = v
	fmt.Println(a.Abs())

	var i interface{} = "hello" // this is like any i can add any type

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)

	fmt.Println(s, ok)

	g, ok := i.(float64)
	fmt.Println(g, ok)

	g = i.(float64) //panic
	fmt.Println(g)

	// type switch section
	do(s)

	// stringer section

	a3 := Person{"Arthure Dent", 53}
	z3 := Person{"Nithin", 26}

	fmt.Println(a3, z3)
}

//type switch

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%q is !v is %v\n", v, v*2)
	case string:
		fmt.Println(v)
	default:
		fmt.Println("i don't know the type")
	}

}

//stringer

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
