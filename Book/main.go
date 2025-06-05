package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("whats your fav book")
	name, _ := reader.ReadString('\n') //includes newline character
	fmt.Println("whats the author's name", name)
	// fmt.Println(reflect.TypeOf(name))
	fmt.Println("what edition is this book")
	favNum, _ := reader.ReadString('\n') //includes newline character
	fmt.Println("Your fav number book is ", favNum)
	fmt.Println("Enter your fun fact")
	funFact, _ := reader.ReadString('\n') //includes newline character
	fmt.Println("Your fun fact is ", funFact)
}
