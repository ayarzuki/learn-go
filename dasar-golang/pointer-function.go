package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesian(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var alamat = Address{
		City:     "Semarang",
		Province: "Jawa Tengah",
		Country:  "",
	}

	var alamatPointer *Address = &alamat
	ChangeCountryToIndonesian(alamatPointer)

	// ChangeCountryToIndonesian(&alamat)
	fmt.Println(alamat)
}
