package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var one Customer
	one.Name = "Mulzahrian One"
	one.Address = "Indonesia"
	one.Age = 21

	fmt.Println(one)

	Mone := Customer{
		Name:    "Ari",
		Address: "Indonesia",
		Age:     22,
	}
	fmt.Println(Mone)

	//Fast way declaration

	rian := Customer{"Rian", "Indonesia", 22}
	fmt.Println(rian)

}
