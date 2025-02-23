package main

import "fmt"

func main(){
	fmt.Println("learing for loop in golang")
	// for i:=0;i<5;i++{
	// 	fmt.Println(i)
	// }

	// for infinite loop with having break statement
	// counter:=0
	// for{
	// 	counter++
	// 	fmt.Println(counter)
	// 	if counter ==3{
	// 		break
	// 	}
	// }

	numbers := []int{1,2,3,4,5}

	for index,value := range numbers{
		fmt.Printf("index is %d and value is %d \n",index,value)
	}

	name := "Harsh Maniar"
	for index,char := range name{
		fmt.Printf("Index is %d and char is %c\n",index,char)
	}
}