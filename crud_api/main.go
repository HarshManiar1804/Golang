package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type TODO struct{
	UserId int64 `json:"userId"`
	Id int64 `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func performGetRequest(){
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
func performPostRequest(){
	myURL := "https://jsonplaceholder.typicode.com/todos/1"
	todo:= TODO{
		UserId :23,
		Id :201,
		Title :"newTitle",
		Completed :false,
	}

	// convert todo structure into json format
	todoJsonData,_:=json.Marshal(todo)

	// convert json data to string
	jsonString := string(todoJsonData)

	// convert string data into reader
	jsonReader:=strings.NewReader(jsonString)

	// send POST request
	res,err:=http.Post(myURL,"application/json",jsonReader)
	if err != nil{
		fmt.Println("posr req err")
	}

	defer res.Body.Close()
	
	data,_ :=io.ReadAll(res.Body)

	fmt.Println("response is",string(data))
	fmt.Println("response status",res.StatusCode)
}

func performUpdateRequest(){
	myURL2 := "https://jsonplaceholder.typicode.com/todos/1"
	todo:= TODO{
		UserId :234,
		Id :202,
		Title :"updated newTitle ",
		Completed :true,
	}

	// convert todo structure into json format
	todoJsonData,_:=json.Marshal(todo)
	
	// convert json data to string
	jsonString := string(todoJsonData)

	// convert string data into reader
	jsonReader:=strings.NewReader(jsonString)

	const myURL = "https://jsonplaceholder.typicode.com/todos/1"

	req,_:= http.NewRequest(http.MethodPut,myURL2,jsonReader)

	req.Header.Set("Content-type", "application/json")


	// send the request
	client := http.Client{}
	res,_:=client.Do(req)
	defer res.Body.Close()


	// read the response

	data,_:=ioutil.ReadAll(res.Body)

	fmt.Println("response is",string(data))
	fmt.Println("response status",res.StatusCode)
}

func performDeleteRequest(){

	const myURL = "https://jsonplaceholder.typicode.com/todos/1"

	req,_:= http.NewRequest(http.MethodDelete,myURL,nil)

	req.Header.Set("Content-type", "application/json")

	// send the request
	client := http.Client{}
	res,_:=client.Do(req)
	defer res.Body.Close()


	fmt.Println("response status",res.StatusCode)
}
func main(){
	// performGetRequest()
	//  performPostRequest()
	//  performUpdateRequest()
	performDeleteRequest()
	
}