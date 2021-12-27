package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address := Address{"Marline", "Singapore", ""}
	changeAddressToIndonesia(&address)

	fmt.Println(address)
}

func changeAddressToIndonesia(address *Address) {
	address.Country = "Singapore"
}
