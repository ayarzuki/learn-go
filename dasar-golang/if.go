package main

import "fmt"

func main () {
	var name = "Rully"

	if name == "Eko" {
		fmt.Println("hAI")
	} else if name == "Rully" {
		fmt.Println("hAII")
	} else {
		fmt.Println("hAIII")
	}

	if length := len(name); length > 5 {
		fmt.Println("name too long")
	} else {
		fmt.Println("name too short")
	}
}