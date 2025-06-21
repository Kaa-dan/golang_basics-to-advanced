package main

import (
	"errors"
	"fmt"
)

const LoginToken string = "gghdd" // public variable starts with capital leter

func main() {
	var username string = "nithin raj"

	fmt.Println(username)

	fmt.Printf("Variable is of type : %T \n", username)

	//boolean

	var isLoggedInd bool = false

	fmt.Println(isLoggedInd)

	fmt.Printf("Variable is of type : %T \n", isLoggedInd)

	// uint

	var smallvalue uint8 = 255

	fmt.Println(smallvalue)

	fmt.Printf("Variable is of type : %T \n", smallvalue)

	//default values and some aliases
	var anothervalue int

	fmt.Println(anothervalue)

	fmt.Printf("Variable is of type : %T \n", anothervalue)

	// implicit type

	var websit = "nithin.com"

	fmt.Println(websit)

	// no var style

	numberOfUser := 3000

	fmt.Println(numberOfUser)

	// errors

	err := errors.New("hello")
	fmt.Println(err)
}
