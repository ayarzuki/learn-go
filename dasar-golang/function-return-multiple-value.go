package main

import "fmt"

func getHello() (string, string, string, string) {
	return "Hello", "World", "and", "Surabaya"
}

func main() {
	firstName, secondName, thirdName, lastName := getHello()
	fmt.Println(firstName, secondName, thirdName, lastName)

	// Menghiraukan multiple return value
	newName, _, _, newestName := getHello()
	fmt.Println(newName, newestName)
}