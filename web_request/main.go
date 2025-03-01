package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	// fetching data using get call
	res,err:=http.Get("https://jsonplaceholder.typicode.com/todos/1")
	// error handling 
	if err != nil{
		fmt.Println("error while fetching api", err)
		return
	}
	// close the Res.Body object
	defer res.Body.Close()

	// read the response
	data,err := ioutil.ReadAll(res.Body)
	
	// error handling 
	if err != nil{
		fmt.Println("error while fetching api", err)
		return
	}
	fmt.Println("response", string(data))
}