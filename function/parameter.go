package main

import "fmt"

func main() {
	sayHello("Ari", "Rian")
}

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
