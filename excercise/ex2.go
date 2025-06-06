package main

import "fmt"

func ex2(){
	var Temperature int
	var Humidity int
	
	fmt.Println("whats the Temperature?")
	fmt.Scanln(&Temperature)

	fmt.Println("whats the Humidity?")
	fmt.Scanln(&Humidity)

	if (Temperature <10){
    if (Humidity>80){
			fmt.Println("Wear a waterproof coat.")
		}else{
			fmt.Println("wear a warm coat. ")
		}
}else if (Temperature >=10 && Temperature <=20){
	if (Humidity>80){
		fmt.Println("Wear a waterfproof jacket.")
	}else{
		fmt.Println("Wear a jacket. ")
	}
}else{
if (Humidity>80){
	fmt.Println("Wear light clothes and carry an umbrella.")
	}else{
		fmt.Println("Wear light clothes.")
	}
}	
}