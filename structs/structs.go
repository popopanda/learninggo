//allows the construction of custom data types

package main

import "fmt"

func main() {
	type books struct {
		Author string
		Title  string
		bookID int
	}
	// initializes variables with type of books with 0 values
	// var book1 books
	//or
	book2 := new(books) //gives a pointer

	book2.Author = "Becky Doe"
	book2.Title = "Learning how to program with Golang"
	book2.bookID = 20

	book3 := books{
		Author: "John Doe",
		Title:  "How to train shiba inus",
		bookID: 1,
	}

	fmt.Println(book2)

	book3.Author = book2.Author
	fmt.Printf("\nThe author of the book is: %v\nThe title of the book is: %v\nThe book ID is: %v\n", book3.Author, book3.Title, book3.bookID)

}
