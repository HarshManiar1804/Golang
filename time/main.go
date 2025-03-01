package main

import (
	"fmt"
	"time"
)

func main(){
	currentTime := time.Now()
	fmt.Println("Current time is :-", currentTime)
	formattedTime := currentTime.Format("02-01-2006 , 15:04:05")
	fmt.Println("formatted time is :-", formattedTime)

	layout_str:="2006-01-02"
	dateStr :="2024-04-18"
	formatted_dateStr,_ := time.Parse(layout_str,dateStr)
	fmt.Println("time parse: ",formatted_dateStr)
}