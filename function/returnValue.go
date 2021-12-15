package main

import "fmt"

func main() {
	result := getHello("Ari")
	fmt.Println(result)
	secondResult := Hi("Algorithm")
	fmt.Println(secondResult)
}

func Hi(subject string) string {
	if subject == "" {
		return "No Subject"
	} else {
		return "Ur Subject is " + subject
	}
}

func getHello(name string) string {
	return "Hellow" + " " + name
}
