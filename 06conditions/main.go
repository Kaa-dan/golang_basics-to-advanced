package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {

	if v := math.Pow(1, 2); v < 10 {
		fmt.Println(v)
	} else {
		fmt.Println("Kannas kadalas")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS x.")
	default:
		fmt.Printf("%s.\n", os)
	}

	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tommorow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("to far away")
	}

	//switch like else if

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good evening")
	}

}
