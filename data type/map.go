package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Rian",
		"address": "Jakarta",
	}

	person["title"] = "Programmer"

	fmt.Println(person["title"])

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	fmt.Println(len(person))

	book := make(map[string]string)
	book["title"] = "Go-Lang Book"
	book["author"] = "Mulzahrian"
	book["wrong"] = "Ups"

	fmt.Println(book)

	delete(book, "wrong")

	fmt.Println(book)
}
