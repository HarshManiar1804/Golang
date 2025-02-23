package main

import (
	"fmt"
	"strconv"
)

func main(){
	var num int = 18
	var num2 float64 = 3.14
	var num3 int = int(num2)
	var num4 float64 = float64(num)
	var stringVar string = strconv.Itoa(num)
	fmt.Printf("num: %T\n", num)
	fmt.Printf("num2: %T\n", num2)
	fmt.Printf("num3: %T\n", num3)
	fmt.Printf("num4: %T\n", num4)
	fmt.Printf("stringVar: %T\n", stringVar)

	var str string = "18"
	strInt,_:= strconv.Atoi(str)

	fmt.Printf("strInt: %T\n", strInt)
}