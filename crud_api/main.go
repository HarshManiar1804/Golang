package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TODO struct{
	UserId int64 `json:"userId"`
	Id int64 `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`

}
func main(){
	res,err:=http.Get("https://jsonplaceholder.typicode.com/todos/1")
	// error handling 
	if err != nil{
		fmt.Println("error while fetching api", err)
		return
	}
	// close the Res.Body object
	defer res.Body.Close()

	// read the response
	var todo TODO
	_= json.NewDecoder(res.Body).Decode(&todo)

	fmt.Println("data of first todo is ",todo)

}