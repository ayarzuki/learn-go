package main

import (
	"fmt"
	"learn-go/dasar-golang/database"
	// _ "learn-go/dasar-golang/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}