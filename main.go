package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")
	foo()
	fmt.Println("SOMETHING MORE")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("FOO.")
}

func bar() {
	fmt.Println("BAR.")
}