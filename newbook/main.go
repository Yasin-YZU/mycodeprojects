package main

import (
	"bufio"
	"fmt"
	"os"
)

type Book struct{
	Authorname string
	Title string
	Price int
}

func main(){

	reader := bufio.NewReader(os.Stdin)


	var newBook [3]Book
	for i := 0; i < 3; i++{
		fmt.Printf("\nEnter details for Book #%d\n", i+4)

		fmt.Print("Enter bookname: ")
		name, _ := reader.ReadString('\n')
		newBook[i.na]
	}
	//book1 := Book{
  //  Authorname:"Roald Dahl",
	//	Title: "Big freindly giant",
	//	Price: 8,
//	}
//	book2 := Book{
//    Authorname:"Bob the builder",
//		Title: "Building blocks",
//		Price: 7,
//	}

//  fmt.Println("Author information:")
//  fmt.Println("Authorname:book.Authorname")
// fmt.Println("Title:",book1.Title)
//  fmt.Println("Price:",book1.Price)

//	fmt.Println("Author information:")
//  fmt.Println("Authorname:book.Authorname")
//  fmt.Println("Title:",book2.Title)
//  fmt.Println("Price:",book2.Price)
}



