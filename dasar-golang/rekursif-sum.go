package main

import "fmt"

func rekursifSum(value int) int {
	if value == 0 {
		return 0
	} else {
		return value + rekursifSum(value-1)
	}
}

func main() {
	hasil := rekursifSum(10)
	fmt.Println(hasil)
}