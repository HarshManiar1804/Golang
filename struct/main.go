package main

import "fmt"

type Person struct{
	Name string
	Age int
}

type Contact struct{
	Phone string
	Email string
}
type Address struct{
	HouseNo int
	Area string
}

type Employee struct{
	Person_Details Person
	Contact_Details Contact
	Address_Details Address
}
func main(){
	// 1st method
	
	// var p1 Person
	// p1.Name = "Harsh"
	// p1.Age = 21
	

	// 2nd method
	
	// p2 :=Person{
	// 	Name :"Meet",
	// 	Age :22,
	// }

	 // 3rd method
	
	 // var p3 =  new(Person)
	// p3.Name = "Parth"
	// p3.Age = 23


	// fmt.Printf("Name is %s and Age is %d\n",p1.Name,p1.Age)
	// fmt.Printf("Name is %s and Age is %d\n",p2.Name,p2.Age)
	// fmt.Printf("Name is %s and Age is %d\n",p3.Name,p3.Age)


	var employee Employee

	employee.Person_Details.Name = "Harsh"
	employee.Person_Details.Age = 21

	employee.Contact_Details.Email = "harsh@gmail.com"
	employee.Contact_Details.Phone = "1234567890"

	employee.Address_Details.HouseNo = 18
	employee.Address_Details.Area = "Darbargadh"

	fmt.Println("Employee details is ", employee)
}

