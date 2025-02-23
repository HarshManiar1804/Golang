package main

import "fmt"

func main(){
	fmt.Println("learning switch")

	room_no := 4

	switch room_no{
	case 1,4,5:
		fmt.Println("room no 1 or 4 or 5")
	case 2:
		fmt.Println("room no 2")
	case 3:
		fmt.Println("room no 3")
	default:
		fmt.Println("other rooms")
	}

	// we can put simple conditions also 
	temperature := 28
	switch {
	case temperature < 15:
		fmt.Println("cold day")
	case temperature > 15 && temperature <=30:
		fmt.Println("sunny day")
	case temperature > 30:
		fmt.Println("hot day")
	}
}