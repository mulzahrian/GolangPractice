package main

import "fmt"

func main() {
	name := "Rian"

	if name == "Ari" {
		fmt.Println("Hellow Ari")
	} else {
		fmt.Println("Hooo, Rian is ur Name")
	}
	//Another version

	Number := "1"

	if Number == "1" {
		fmt.Println("ur the one")
	} else if Number == "2" {
		fmt.Println("ur the second one")
	} else {
		fmt.Println("i dunno who are u")
	}

	//Short statement
	if length := len(name); length > 5 {
		fmt.Println("Ur name to long")
	} else {
		fmt.Println("Ur name was right")
	}
}
