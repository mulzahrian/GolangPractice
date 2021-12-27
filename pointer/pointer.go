package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Rumbai", "Pekanbaru", "Indonesia"}
	//when u add & the data u change on address2 will automaticly replace to address1
	address2 := &address1

	address2.City = "Bengkalis"

	//address2 = &Address{"Jakarta", "JakBar", "Indonesia"}
	*address2 = Address{"Jakarta", "JakBar", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	var address3 *Address = new(Address)
	var address4 *Address = new(Address)
	address4.City = "Batam"
	fmt.Println(address3)
	fmt.Println(address4)

}
