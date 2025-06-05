package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your full name")
	name, _ := reader.ReadString('\n') //includes newline character
	fmt.Println("Your name is ", name)
	// fmt.Println(reflect.TypeOf(name))
	fmt.Print("Enter your favourite number")
	favNum, _ := reader.ReadString('\n') //includes newline character
	fmt.Println("Your fav number is ", favNum)
	fmt.Print("Enter your fun fact")
	funFact, _ := reader.ReadString('\n') //includes newline character
	fmt.Println("Your fun fact is ", funFact)
}
