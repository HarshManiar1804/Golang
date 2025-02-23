package main

import "fmt"

func main(){
	fmt.Println("learning array in go lang")

	// two type of array initialization
	var name[5] string
	name[0] = "harsh"
	name[1] = "ruhi"
	name[2] = "shruvi"
	name[3] = "jay"
	name[4] = "kinar"

	fmt.Printf("value of the %s\n",name[1] )


	var rollNo = [5]int{1,2,3,4,5}
	fmt.Printf("value of the %d\n",rollNo[1] )
	// length of an array
	fmt.Println("Length of the Name array is :",len(name))

	// in array the default value is set as 0 whenever we create the array type of int 
	// similarly with every datatype this do the same 
	var marks[5]int 
	fmt.Println("marks of all the students are :" , marks)

}