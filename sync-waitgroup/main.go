package main

import (
	"fmt"
	"sync"
)

func worker(i int , wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("Worker %d working\n",i)
	// some tasks are happing
	fmt.Printf("Worker %d ends\n",i)
}
func main(){

	var wg sync.WaitGroup
	for i:=1;i<=3;i++{
		wg.Add(1) // increment the WaitGroup counter
		go worker(i , &wg)
	}
	// wait for all workers to complete their tasks 
	wg.Wait()
	fmt.Println("Workers tasks completed")
}