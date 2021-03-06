package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := &address1
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	// var address4 *Address = &Address{"Surabaya", "Jawa Timur", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Bandung"

	// address2 = &Address{"Jakarta", "Jawa Barat", "Indonesia"}
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	// address4.City = ""
	address4.City = "Jogjakarta"
	fmt.Println(address4)
}