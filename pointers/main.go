package main

import "fmt"

func modifyNumber(num *int){
	*num += 10
}

func main(){
	// pointer will store the address of the variable in memory
	var num int = 2
	var ptr *int = &num

	fmt.Println("num: ", num)
	fmt.Println("ptr: ", *ptr)
	fmt.Println("ptr address: ", ptr)

	var pointer *int
	// default value of pointer is nil
	if(pointer == nil){
		fmt.Println("pointer is nil")
	}

	modifyNumber(&num)
	fmt.Println("num after modifyNumber: ", num)
}