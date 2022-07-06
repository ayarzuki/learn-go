package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello "
	} else {
		return "Hello " + name
	}
}

func main() {
	result := "Budi"
	fmt.Println(getHello(result))

	fmt.Println(getHello("Tita"))

	fmt.Println(getHello(""))
}