package main

import "fmt"

func main(){
	fmt.Println("Learing conditions")
	age:= 18

	if age > 18{
		fmt.Println("age is greater then 18")
	}else if age < 18{
		fmt.Println("age is less then 18")
	}else{
		fmt.Println("age is equal to 18")
	}
}