package main

import "fmt"

func main() {
	type noKTP string
	type Married bool

	var noKTPEko noKTP = "1232323"
	var marredstatus Married = false

	fmt.Println(noKTPEko)
	fmt.Println(marredstatus)
}