package main

import "fmt"

func main(){
	// name -> grade
	studentGrades := make(map[string]int)

	studentGrades["harsh"] = 90
	studentGrades["meet"] = 99
	studentGrades["parth"] = 95


	studentSurname := map[string]string{
		"harsh":"maniar",
		"meet":"trivedi",
		"parth":"hajari",
	}
	// fmt.Println("Marks of meet is:- ", studentGrades["meet"])

	// studentGrades["meet"] = 97

	// fmt.Println("New Marks of meet is:- ", studentGrades["meet"])

	// delete(studentGrades,"meet")

	// grade , exists := studentGrades["harsh"]
	// fmt.Printf("grade is %d and if the value is exists ? -> %t\n ", grade,exists)

	// fmt.Println("after delete Marks of meet is:- ", studentGrades["meet"])
	for index,value := range studentGrades{
		fmt.Printf("index is %s and value is %d\n",index,value)
	}
	fmt.Println("-----------------")
	for index,value := range studentSurname{
		fmt.Printf("name is %s and surname is %s\n",index,value)
	}
}