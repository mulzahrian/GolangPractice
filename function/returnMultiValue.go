package main

import "fmt"

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

}

func getFullName() (string, string) {
	return "Mulzahrian", "One"
}
