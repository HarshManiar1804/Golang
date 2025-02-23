package main

import "fmt"

func simpleFunction(){
	fmt.Println("This is a simple function")
}

// syntax for function is 

// func functionName (parameter1 type, parameter2 type, parameter3 type) (returnType)
// {
//	 function body
// }

// if we have same data type for all the parameters then we can write like this (the common data type needs to be written only once after the last parameter)

// func addFunction(a,b,c int)

func addFunction(a int , b int , c int )int{
	return a + b + c
}

func subFunction (a , b int)(result int){
	result = a-b
	return 
}
func main() {

	fmt.Println("This is the main function")
	simpleFunction()
	resultAdd := addFunction(1,2,3);
	resultSub := subFunction(10,5)
	fmt.Printf("resule of the add function is %d\n",resultAdd)
	fmt.Printf("resule of the sub function is %d\n",resultSub)
}