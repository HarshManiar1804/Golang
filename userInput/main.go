package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// age := 0
	// name := ""
	// height := 0.0
	fmt.Print("Enter your name: ")
	// fmt.Scan(&name)
	// fmt.Print("Enter your age: ")
	// fmt.Scan(&age)
	// fmt.Print("Enter your height: ")
	// fmt.Scan(&height)
	// fmt.Printf("My age is %d and name is %s and height is %.2f",age,name, height)

	reader := bufio.NewReader(os.Stdin)
	name , _ := reader.ReadString('\n')
	fmt.Printf("My name is %s",name)
}