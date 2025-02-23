package main

import "fmt"

func divide(a float64 , b float64 )(float64 , error){

	if (b == 0){
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return a/b,nil
}
func main(){

	resultDivide , err := divide(10,0);
	if ( err != nil){	
		fmt.Println("error handling")
	}
	fmt.Printf("result of the divide is %.2f \n",resultDivide, )
	// resultDivide,_ := divide(10,0)
	// fmt.Printf("result of the divide is %.2f \n",resultDivide, )
}