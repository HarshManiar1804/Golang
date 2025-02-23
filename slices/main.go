package main

import "fmt"

// slice is a flexible and dynamic data structure which is more powerful alternative to arrays.
// it is just like vector in c++ stl library
func main(){
	fmt.Println("Learning slices in golang")
	numbers := []int{1,2,3,4,5}
	numbers = append(numbers, 6,7,8,9,0)
	fmt.Println("value in numbers are", numbers)
	fmt.Printf("length of the numbers is %d\n", len(numbers))
	fmt.Printf("data type of numbers is %T\n", numbers)
	fmt.Printf("length of the numbers is %d\n", len(numbers))


	name := []string{}
	name = append(name, "harsh","ruhi")
	fmt.Println("value of names are ",name)


	age := make([]int,3,5)
	age = append(age, 32)
	fmt.Println("value of age are ",age)
	fmt.Println("len of age are ",len(age))
	fmt.Println("cap of age are ",cap(age))

	// if you want to initialize the slice without providing the values then you need to must use the "make" function

}