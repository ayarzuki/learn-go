package main

import "fmt"

func logging() {
	fmt.Println("Selesai panggil function")
}

func runApplication(value int) {
	defer logging()
	// logging()
	fmt.Println("Panggil function")
	result := 10 / value
	fmt.Println("result", result)
}

func main() {
	runApplication(0)
}