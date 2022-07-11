package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (a Customer) sayHuu(){
	fmt.Println("Huuuu, from", a.Name)
}

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Jalan Raya"
	eko.Age = 25

	eko.sayHi("Tita")
	eko.sayHuu()
}