package main

import "fmt"

type Filter func(string) string

func main() {
	sayHelloWithFilter("One", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("dog", filter)
}
func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("hello ", filter(nameFiltered))
}

func spamFilter(name string) string {
	if name == "dog" {
		return "..."
	} else {
		return name
	}
}
