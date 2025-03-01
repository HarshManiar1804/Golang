package main

import "fmt"

func add(a,b int ) int{
	return a+b
}

func main(){

	fmt.Println("start")
	ans:=add(5,10)
	defer fmt.Println("ans is", ans)
	defer fmt.Println("middle")
	fmt.Println("end")


}