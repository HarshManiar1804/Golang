package main

import (
	"fmt"
	"strings"
)

func main(){

	var1 := "Hello,Hii,Hola,Hii,Hello,Hii"
	// parts_var1 := strings.Split(var1,",")
	// for index,parparts_var1 := range parts_var1{
	// 	fmt.Println(index,parparts_var1)
	// }
	count :=strings.Count(var1,"Hii")
	fmt.Printf("Count of Hii is %d\n",count)

	str := "    Hello, World!    "
	trimmed := strings.TrimSpace(str)
	fmt.Println(trimmed)

	str1:="harsh"
	str2:="nishith"
	str3:="maniar"
	joinedStr := strings.Join(([]string{str1,str2,str3})," ")
	fmt.Println(joinedStr)

}