package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Jalan Raya"
	eko.Age = 25

	fmt.Println(eko)
	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	joko := Customer{
		Name: "Joko",
		Address: "Jalan Jambu",
		Age: 30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Jakarta", 35}
	fmt.Println(budi)
}
