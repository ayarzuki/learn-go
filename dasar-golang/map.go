package main

import "fmt"

func main () {
	person := map[string]string{
		"name": "John",
		"address": "Subang",
	}

	person["title"] = "IT"
	fmt.Println(person)
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Book"
	book["author"] = "John"
	book["page"] = "ups"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "page")

	fmt.Println(book)
	fmt.Println(len(book))
}