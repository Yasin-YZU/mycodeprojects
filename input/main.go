package main

import "fmt"

func main(){
var name string
var age int
var fact string 

fmt.Println("What's your name?")
fmt.Scanln(&name)

fmt.Println("What's your age?")
fmt.Scanln(&age)

fmt.Println("What's the fact?")
fmt.Scanln(&fact)

fmt.Println("My name is: ", name)
fmt.Println("My age is: ", age)
fmt.Println("My fact is: ", fact)
}