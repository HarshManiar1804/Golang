package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	/*
	fmt.Println("file in golang")

	file , err:= os.Create("example.txt")

	if err != nil{
		fmt.Println("err in file")
		return 
	}
	content :="hello world"

	byte,error := io.WriteString(file,content + "\n")
	if error != nil{
		fmt.Println("err in file",error)
		return 
	}
	fmt.Println("byte in file",byte)

	defer file.Close()
	*/

	file,err :=os.Open("example.txt")
	if err != nil{
		fmt.Println("err in file")
		return 
	}
	defer file.Close()


	// creating a buffer to read the file content

	buffer := make([]byte,1024)

	// read the file content into the buffer

	for{
		n,err :=file.Read(buffer)
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Println("Error while reading the file",err)
		}
		// process to read the content
		fmt.Println(string(buffer[:n]))
	}

	// read the entire file into one byte size

	content,err := os.ReadFile("example.txt")
	if err != nil{
		fmt.Println("Error while reading the file",err)
	}
	fmt.Println("entire file is ",string(content))
}