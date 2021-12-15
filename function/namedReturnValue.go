package main

import "fmt"

func main() {
	firstName, middleName, lastName := getComplateName()
	fmt.Println(firstName, middleName, lastName)
}

func getComplateName() (firstName, middleName, lastName string) {
	firstName = "Rian"
	middleName = "One"
	lastName = "Ari"

	return firstName, middleName, lastName
}
