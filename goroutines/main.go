package main

import (
	"fmt"
	"time"
)

func sayHello(){
	fmt.Println("say hello")
	// time.Sleep(1000 * time.Millisecond)
	fmt.Println("say hello after wait")

}
func sayHi(){
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("say hii")
}

func main(){
	// goroutines is a key feature of the GO lang that allow us to run functions concurrently or you can say parellel other parts of the program.
	
	// start goroutines
	go sayHello()
	go sayHi()
	time.Sleep(900 * time.Millisecond)
}