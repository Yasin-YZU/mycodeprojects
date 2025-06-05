package main

import "fmt"

func main(){
	for aborder := 'A'; aborder <= 'Z'; aborder++{
		fmt.Printf("%c and type: %T\n", aborder, aborder)
	}
}