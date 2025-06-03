package main

import "fmt"

func main(){
	var examResult int
	var ifPassed bool

	fmt.Println("Whats your exam results?")
	fmt.Scanln(&examResult)
	
	
	if( examResult > 50){
		ifPassed=true
	}else{
		ifPassed=false
	}

	fmt.Println("Passed: ", ifPassed)
}