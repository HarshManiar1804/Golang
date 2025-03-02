package main

import (
	"encoding/json"
	"fmt"
)

type Person struct{
	Name string	`json:"name"`
	Phone_no string	`json:"phoneNo"`
	Age int	`json:"age"`
}
func main(){

	person := Person{
		Name: "Harsh",
		Phone_no: "0987654321",
		Age: 22,
	}
	fmt.Println("Person data is ",person)

	// convert person into JSON encoding (Marshalling)
	personJson,_ :=json.Marshal(person)
	fmt.Println("---------------------------")
	fmt.Println("Person data is ",string(personJson))

	// Decoding (Unmarshalling)
	var personData Person 
	err :=json.Unmarshal(personJson,&personData)
	if err != nil{
		fmt.Println("Error while unmarshal the data")
		return
	}
	fmt.Println("---------------------------")
	fmt.Println("Unmarshal data", personData)

}